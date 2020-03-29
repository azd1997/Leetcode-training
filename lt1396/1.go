package lt1396

// 设计地铁系统

type PersonKey struct {
	id  int
	tin int
}

type PersonInfo struct {
	id   int
	in   string
	tin  int
	out  string // ""表示还未出站
	tout int
}

type StationInOut struct { // 入站、出站组合
	in  string
	out string
}

type StationTotal struct { // 入站、出站组合
	totalDuration int // 总时长
	totalNum      int // 总次数
}

type UndergroundSystem struct {
	personTable  map[PersonKey]*PersonInfo
	personLatest map[int]*PersonKey // <id, PersonKey> 每个人最近一次的入站
	stations     map[StationInOut]*StationTotal
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		personTable:  make(map[PersonKey]*PersonInfo),
		personLatest: make(map[int]*PersonKey), // <id, PersonKey> 每个人最近一次的入站
		stations:     make(map[StationInOut]*StationTotal),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {

	//fmt.Println(id, stationName, t, this)

	k := PersonKey{id, t}
	info := &PersonInfo{id: id, in: stationName, tin: t}
	if v, ok := this.personTable[k]; ok && v != nil {
		return // 说明有问题
	} else {
		this.personTable[k] = info
		this.personLatest[id] = &k
	}

}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	k := this.personLatest[id]
	var info *PersonInfo
	var ok bool
	if info, ok = this.personTable[*k]; !ok || (ok && info.out != "") {
		return // 没有匹配的出站信息
	}
	info.out = stationName
	info.tout = t

	// 需要更新入出站信息
	sio := StationInOut{in: info.in, out: info.out}
	if st, ok := this.stations[sio]; !ok {

		this.stations[sio] = &StationTotal{info.tout - info.tin, 1}
		//       fmt.Println("st1:",sio, this.stations[sio])
	} else {
		st.totalDuration += info.tout - info.tin
		st.totalNum++
		//  fmt.Println("st:", st)
	}

	// 把个人入站出站信息清除
	this.personTable[*k] = nil
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	k := StationInOut{in: startStation, out: endStation}
	total := this.stations[k]
	//fmt.Println(k, total)
	if total == nil {
		return 0
	}
	return float64(total.totalDuration) / float64(total.totalNum)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
