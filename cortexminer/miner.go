package cortexminer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	//	"github.com/ethereum/go-ethereum/PoolMiner/miner/libcuckoo"
	"github.com/ethereum/go-ethereum/PoolMiner/config"

	"plugin"
	"strconv"
)

func checkError(err error, func_name string) {
	if err != nil {
		log.Println(func_name, err.Error())
		//		os.Exit(1)
	}
}

func (cm *Cortex) read() map[string]interface{} {
	rep := make([]byte, 0, 1024) // big buffer
	for {
		tmp, isPrefix, err := cm.reader.ReadLine()
		if err == io.EOF {
			log.Println("Tcp disconnectted")
			cm.consta.lock.Lock()
			defer cm.consta.lock.Unlock()
			cm.conn.Close()
			cm.conn = nil
			cm.consta.state = false
			return nil
		}
		checkError(err, "read()")
		rep = append(rep, tmp...)
		if isPrefix == false {
			break
		}
	}
	// fmt.Println("received ", len(rep), " bytes: ", string(rep), "\n")
	var repObj map[string]interface{}
	err := json.Unmarshal(rep, &repObj)
	checkError(err, "read()")
	return repObj
}

func (cm *Cortex) write(reqObj ReqObj) {
	req, err := json.Marshal(reqObj)
	checkError(err, "write()")

	req = append(req, uint8('\n'))
	_, _ = cm.conn.Write(req)
}

//	init cortex miner
func (cm *Cortex) init() *net.TCPConn {
	log.Println("Cortex Init")
	cm.consta.lock.Lock()
	defer cm.consta.lock.Unlock()
	//cm.server = "cortex.waterhole.xyz:8008"
	//cm.server = "localhost:8009"
	//cm.account = "0xc3d7a1ef810983847510542edfd5bc5551a6321c"
	tcpAddr, err := net.ResolveTCPAddr("tcp", cm.param.Server)
	checkError(err, "init()")

	cm.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	cm.conn.SetKeepAlive(true)
	cm.conn.SetNoDelay(true)
	log.Println("Cortex connect successfully")
	checkError(err, "init()")
	cm.consta.state = true
	cm.reader = bufio.NewReader(cm.conn)
	log.Println("Cortex Init successfully")
	return cm.conn
}

//	miner login to mining pool
func (cm *Cortex) login() {
	log.Println("Cortex login ...")
	var reqLogin = ReqObj{
		Id:      73,
		Jsonrpc: "2.0",
		Method:  "ctxc_submitLogin",
		Params:  []string{cm.param.Account},
	}
	cm.write(reqLogin)
	cm.read()
	log.Println("Cortex login suc")
}

//	get mining task
func (cm *Cortex) getWork() {
	req := ReqObj{
		Id:      100,
		Jsonrpc: "2.0",
		Method:  "ctxc_getWork",
		Params:  []string{""},
	}
	cm.write(req)
}

//	submit task
func (cm *Cortex) submit(sol config.Task) {
	var reqSubmit = ReqObj{
		Id:      73,
		Jsonrpc: "2.0",
		Method:  "ctxc_submitWork",
		Params:  []string{sol.Nonce, sol.Header, sol.Solution},
	}
	cm.write(reqSubmit)
}

var minerPlugin *plugin.Plugin

const PLUGIN_PATH string = "plugins/"
const PLUGIN_POST_FIX string = "_helper.so"

//	cortex mining
func (cm *Cortex) Mining() {
	var iDeviceIds []uint32
	for i := 0; i < len(cm.deviceInfos); i++ {
		iDeviceIds = append(iDeviceIds, cm.deviceInfos[i].DeviceId)
	}

	var minerName string = ""
	if cm.param.Cpu == true {
		minerName = "cpu"
	} else if cm.param.Cuda == true {
		minerName = "cuda"
	} else if cm.param.Opencl == true {
		minerName = "opencl"
	} else {
		os.Exit(0)
	}

	var err error
	minerPlugin, err = plugin.Open(PLUGIN_PATH + minerName + PLUGIN_POST_FIX)
	m, err := minerPlugin.Lookup("CuckooInitialize")
	if err != nil {
		panic(err)
	}
	m.(func([]uint32, uint32, config.Param))(iDeviceIds, (uint32)(len(iDeviceIds)), cm.param)
	go func() {
		for {
			cm.printHashRate()
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		for {
			//cm.consta.lock.Lock()
			//defer cm.consta.lock.Unlock()
			consta := cm.consta.state
			if consta == false {
				cm.init()
				cm.login()
			} else {
				break
				//return
			}
		}
		time.Sleep(1 * time.Second)
		cm.miningOnce()
	}
}

func (cm *Cortex) printHashRate() {
	var devCount = len(cm.deviceInfos)
	var fanSpeeds []uint32
	var temperatures []uint32
	m, err := minerPlugin.Lookup("Monitor")
	if err != nil {
		panic(err)
	}
	fanSpeeds, temperatures = m.(func(uint32) ([]uint32, []uint32))(uint32(devCount))
	var total_solutions int64 = 0
	for dev := 0; dev < devCount; dev++ {
		var dev_id = cm.deviceInfos[dev].DeviceId
		gps := (float32(1000.0*cm.deviceInfos[dev].Gps) / float32(cm.deviceInfos[dev].Use_time))
		if cm.deviceInfos[dev].Use_time > 0 && cm.deviceInfos[dev].Solution_count > 0 {
			cm.deviceInfos[dev].Hash_rate = (float32(1000.0*cm.deviceInfos[dev].Solution_count) / float32(cm.deviceInfos[dev].Use_time))
			log.Println(fmt.Sprintf("\033[0;%dmGPU%d GPS=%.4f, hash rate=%.4f, find solutions:%d, fan=%d%%, t=%dC\033[0m", 32+(dev%2*2), dev_id, gps, cm.deviceInfos[dev].Hash_rate, cm.deviceInfos[dev].Solution_count, fanSpeeds[dev], temperatures[dev]))
			total_solutions += cm.deviceInfos[dev].Solution_count
		} else {
			log.Println(fmt.Sprintf("\033[0;%dmGPU%d GPS=%.4f, hash rate=Inf, find solutions: 0, fan=%d%%, t=%dC\033[0m", 32+(dev%2*2), dev_id, gps, fanSpeeds[dev], temperatures[dev]))
		}
	}
	log.Println(fmt.Sprintf("\033[0;36mfind total solutions : %d, share accpeted : %d, share rejected : %d\033[0m", total_solutions, cm.share_accepted, cm.share_rejected))
}

func readNonce() (ret []uint64) {
	fi, err := os.Open("nonces.txt")
	if err != nil {
		log.Println("Error:", err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		var strNonce string = string(a)
		nonce, _ := strconv.ParseInt(strNonce, 10, 64)
		ret = append(ret, uint64(nonce))
	}
	return ret
}

func (cm *Cortex) miningOnce() {
	log.Println("mining once")
	var taskHeader, taskNonce, taskDifficulty string
	var THREAD int = (int)(len(cm.deviceInfos))
	rand.Seed(time.Now().UTC().UnixNano())
	solChan := make(chan config.Task, THREAD)
	taskChan := make(chan config.Task, THREAD)

	m, err := minerPlugin.Lookup("RunSolver")
	if err != nil {
		panic(err)
	}
	m.(func(int, []config.DeviceInfo, config.Param, chan config.Task, chan config.Task, bool) (uint32, [][]uint32))(THREAD, cm.deviceInfos, cm.param, taskChan, solChan, cm.consta.state)
	log.Println(".....")
	cm.getWork()

	go func(currentTask_ *config.TaskWrapper) {
		for {
			msg := cm.read()
			if cm.consta.state == false {
				return
			}
			if cm.param.VerboseLevel >= 4 {
				//log.Println("Received: ", msg)
			}
			reqId, _ := msg["id"].(float64)
			result, _ := msg["result"].(bool)
			if uint32(reqId) == 73 {
				if result {
					cm.share_accepted += 1
				} else {
					cm.share_rejected += 1
				}
			}
			if uint32(reqId) == 100 || uint32(reqId) == 0 {
				workInfo, _ := msg["result"].([]interface{})
				if len(workInfo) >= 3 {
					taskHeader, taskNonce, taskDifficulty = workInfo[0].(string), workInfo[1].(string), workInfo[2].(string)
					log.Println("Get Work in task: ", taskHeader, taskDifficulty)
					currentTask_.Lock.Lock()
					currentTask_.TaskQ.Nonce = taskNonce
					currentTask_.TaskQ.Header = taskHeader
					currentTask_.TaskQ.Difficulty = taskDifficulty
					currentTask_.Lock.Unlock()
					for i := 0; i < THREAD; i++ {
						taskChan <- currentTask_.TaskQ
					}
				}
			}
		}
	}(&config.CurrentTask)
	time.Sleep(2 * time.Second)

	for {
		if cm.consta.state == false {
			return
		}
		select {
		case sol := <-solChan:
			//config.CurrentTask.Lock.Lock()
			//defer config.CurrentTask.Lock.Unlock()
			//task := config.CurrentTask.TaskQ
		//	if sol.Header == task.Header {
				cm.submit(sol)
		//	}

		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
