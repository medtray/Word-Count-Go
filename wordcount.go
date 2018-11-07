package main

import (
  "fmt"
  "os"
  "strconv"
  "os/exec"
  "strings"
  "io/ioutil"
  "path"
  "encoding/json"
  "sort"
  "container/list"
  "bufio"
  "hash/fnv"
  "crypto/sha1"
  "sync"
//add new packages as necessary
)

//add new functions and declerations as necessary

func removeDuplicates(elements []int) []int {
   
    encountered := map[int]bool{}
    result := []int{}

    for v := range elements {
        if encountered[elements[v]] == true {
           
        } else {
           
            encountered[elements[v]] = true
            
            result = append(result, elements[v])
        }
    }
   
    return result
}


type Key_Value struct {
	Key   string
	Value int
}


type Work struct {
	mapers_number      int    
	reducers_number   int    
	PathToFile string 
	File_Result string
		
}


type Key_ValueList []Key_Value

func (m Key_ValueList) Len() int           { return len(m) }
func (m Key_ValueList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Key_ValueList) Less(i, j int) bool { return m[i].Value > m[j].Value }



type key_value_string struct {
	Key   string
	Value string
}


type MRR struct {
	work     Work 
	mapers_number    int
	reducers_number int
	file    string
        file_result string
}



func worker_map(wg *sync.WaitGroup, workNumber int, InputFile string,
	reducers_number int) {  
    defer wg.Done()

    CallMap(workNumber, InputFile,reducers_number) 
}


func worker_reduce(wg *sync.WaitGroup, work int, InputFile string, mapers_number int) {  
    defer wg.Done()

    CallReduce(work, InputFile, mapers_number)
}



func CallMap(workNumber int, InputFile string,
	reducers_number int) {
	
	dir1 := path.Dir(InputFile)
	file1 := path.Base(InputFile)
	lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(workNumber))
	
	file, _ := os.Open(lp)
	
	ff,_ := file.Stat()
	
	sz := ff.Size()
	
	b := make([]byte, sz)
	file.Read(b)
	
	file.Close()

	lds := strings.Fields(string(b))

	rs := list.New()
	for _, j := range lds {
		tt := key_value_string{Key: j, Value: "1"}
		rs.PushBack(tt)
	}
	
	
	for r := 0; r < reducers_number; r++ {
		

		dir1 := path.Dir(InputFile)
		file1 := path.Base(InputFile)
		mapFile := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(workNumber))
		dir2 := path.Dir(mapFile)
		file2 := path.Base(mapFile)
		file,_ = os.Create(path.Join(dir2, file2+"_"+strconv.Itoa(r)))
	
		
		yyy := json.NewEncoder(file)
		for o := rs.Front(); o != nil; o = o.Next() {
			tt := o.Value.(key_value_string)
			sha1Bytes := sha1.Sum([]byte(tt.Key))
			h := fnv.New32a()
			h.Write(sha1Bytes[:])
			h.Sum32()
			
			if h.Sum32()%uint32(reducers_number) == uint32(r) {
				 yyy.Encode(&tt)
				
				
			}
		}
		file.Close()
	}



}



func CallReduce(work int, InputFile string, mapers_number int) {
	tts := make(map[string]*list.List)
	for i := 0; i < mapers_number; i++ {
		

		dir1 := path.Dir(InputFile)
		file1 := path.Base(InputFile)
		mapFile := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(i))
		dir2 := path.Dir(mapFile)
		file2 := path.Base(mapFile)
		lp:=path.Join(dir2, file2+"_"+strconv.Itoa(work))
		
		file, error := os.Open(lp)
		if error != nil {
			
		}
		ec := json.NewDecoder(file)
		
		for {
			var tt key_value_string
			error = ec.Decode(&tt)
			if error != nil {
				break
			}
			_, ppp := tts[tt.Key]
			
			if !ppp {
				tts[tt.Key] = list.New()
			}
			tts[tt.Key].PushBack(tt.Value)
			
			
		}
	
		file.Close()
		
	
	}
	
	var kk []string
	for k := range tts {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	dir1 := path.Dir(InputFile)
	file1 := path.Base(InputFile)
	p:=path.Join(dir1, "RED_"+file1+"_"+strconv.Itoa(work))
	
	file,_ := os.Create(p)
	
	yyy := json.NewEncoder(file)
	for _, k := range kk {
		
		
		yyy.Encode(key_value_string{k, strconv.Itoa(tts[k].Len())})
	}
	file.Close()
}










//BONUS
func WordCount_MR_DMP(inFile string, outFile string, mapers_number int, reducers_number int) {
  

}

func WordCount_MR_SMP(inFile string, outFile string, mapers_number int, reducers_number int) {
var wg sync.WaitGroup
work := Work{mapers_number: mapers_number, reducers_number: reducers_number,PathToFile: inFile,File_Result: outFile}
	lll := new(MRR)
	lll.work = work
	lll.mapers_number = work.mapers_number
	lll.reducers_number = work.reducers_number
	lll.file = work.PathToFile
	lll.file_result= work.File_Result
	infile,_ := os.Open(lll.file)
	ff,_ := infile.Stat()
	sz := ff.Size()
	nk := sz / int64(lll.mapers_number)
	nk += 1
	
	dir1 := path.Dir(lll.file)
	file1 := path.Base(lll.file)
	lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(0))
	outfile,_ := os.Create(lp)
	
	wr := bufio.NewWriter(outfile)
	m := 1
	i := 0

	sc := bufio.NewScanner(infile)
	for sc.Scan() {
		if int64(i) > nk*int64(m) {
			wr.Flush()
			outfile.Close()
			dir1 := path.Dir(lll.file)
			file1 := path.Base(lll.file)
			lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(m))
			outfile, _ = os.Create(lp)
			wr = bufio.NewWriter(outfile)
			m += 1
		}
		to_read := sc.Text() + "\n"
		wr.WriteString(to_read)
		i += len(to_read)
	}
	wr.Flush()
	outfile.Close()


	for i := 0; i < lll.mapers_number; i++ {
		
			wg.Add(1)
        go worker_map(&wg, i, lll.file,lll.reducers_number)
			
			}
		
		
	
	wg.Wait()
	
	for i := 0; i < lll.reducers_number; i++ {
		wg.Add(1)
		  
		 go worker_reduce(&wg, i, lll.file,lll.mapers_number)
	}
	
	wg.Wait()
	tts := make(map[string]string)
	for i := 0; i < lll.reducers_number; i++ {
		
		dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		p:=path.Join(dir1, "RED_"+file1+"_"+strconv.Itoa(i))
		
		file, error := os.Open(p)
		if error != nil {
			
		}
		ec := json.NewDecoder(file)
		for {
			var tt key_value_string
			error = ec.Decode(&tt)
			if error != nil {
				break
			}
			tts[tt.Key] = tt.Value
		}
		file.Close()
	}
	var kk []string
	for k := range tts {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	

	p := make(Key_ValueList, len(tts))

	ii := 0
	for k, v := range tts {
		
		aa, _ := strconv.Atoi(v)
		p[ii] = Key_Value{k, aa}
		ii++
	}

	sort.Sort(p)

	var s []int
	var final_string []string
	for _, v := range p {
	s = append(s, v.Value)
	}
	


	result := removeDuplicates(s)
	  

	for kk:=range result{
	
	var add_string []string
	for _, v := range p {
	    
	    if v.Value == result[kk] { 
	      
		add_string = append(add_string, v.Key)
	
	      
	    }
	  }
	sort.Strings(add_string)
	for kkk:=range add_string{
		
	final_string = append(final_string, add_string[kkk])
	}





	}


for i := 0; i < lll.mapers_number; i++ {
		
		
		dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(i))
		os.Remove(lp)
	
		for j := 0; j < lll.reducers_number; j++ {

			dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		mapFile := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(i))
		dir2 := path.Dir(mapFile)
		file2 := path.Base(mapFile)
			
			os.Remove(path.Join(dir2, file2+"_"+strconv.Itoa(j)))
		}
	}



        file,_ := os.Create(lll.file_result)
	
	for kkkk:= range s{
	fmt.Fprint(file, final_string[kkkk]+" "+strconv.Itoa(s[kkkk])+"\n")

	}
		file.Close()


	
        
	for i := 0; i < lll.reducers_number; i++ {
		dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		p:=path.Join(dir1, "RED_"+file1+"_"+strconv.Itoa(i))
	
		
		os.Remove(p)
	}
 


}





func WordCount_MR_S(inFile string, outFile string, mapers_number int, reducers_number int) {

	work := Work{mapers_number: mapers_number, reducers_number: reducers_number,PathToFile: inFile,File_Result: outFile}
	lll := new(MRR)
	lll.work = work
	lll.mapers_number = work.mapers_number
	lll.reducers_number = work.reducers_number
	lll.file = work.PathToFile
	lll.file_result= work.File_Result

	
	infile,_ := os.Open(lll.file)
	
	
	ff,_ := infile.Stat()
	
	sz := ff.Size()
	
	nk := sz / int64(lll.mapers_number)
	nk += 1
	
	dir1 := path.Dir(lll.file)
	file1 := path.Base(lll.file)
	lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(0))
	outfile,_ := os.Create(lp)
	
	wr := bufio.NewWriter(outfile)
	m := 1
	i := 0

	sc := bufio.NewScanner(infile)
	for sc.Scan() {
		if int64(i) > nk*int64(m) {
			wr.Flush()
			outfile.Close()
			dir1 := path.Dir(lll.file)
			file1 := path.Base(lll.file)
			lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(m))
			outfile, _ = os.Create(lp)
			wr = bufio.NewWriter(outfile)
			m += 1
		}
		to_read := sc.Text() + "\n"
		wr.WriteString(to_read)
		i += len(to_read)
	}
	wr.Flush()
	outfile.Close()

	for i := 0; i < lll.mapers_number; i++ {
		
		   CallMap(i, lll.file, lll.reducers_number)
		
		
	}
	
	for i := 0; i < lll.reducers_number; i++ {
		
		   CallReduce(i, lll.file, lll.mapers_number)
	}
	
	tts := make(map[string]string)
	for i := 0; i < lll.reducers_number; i++ {
		
		dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		p:=path.Join(dir1, "RED_"+file1+"_"+strconv.Itoa(i))
		
		file, error := os.Open(p)
		if error != nil {
			
		}
		ec := json.NewDecoder(file)
		for {
			var tt key_value_string
			error = ec.Decode(&tt)
			if error != nil {
				break
			}
			tts[tt.Key] = tt.Value
		}
		file.Close()
	}
	var kk []string
	for k := range tts {
		kk = append(kk, k)
	}
	sort.Strings(kk)
	

	p := make(Key_ValueList, len(tts))

	ii := 0
	for k, v := range tts {
		
		aa, _ := strconv.Atoi(v)
		p[ii] = Key_Value{k, aa}
		ii++
	}

	sort.Sort(p)

	var s []int
	var final_string []string
	for _, v := range p {
	s = append(s, v.Value)
	}
	


	result := removeDuplicates(s)
	  

	for kk:=range result{
	
	var add_string []string
	for _, v := range p {
	    
	    if v.Value == result[kk] { 
	      
		add_string = append(add_string, v.Key)
	
	      
	    }
	  }
	sort.Strings(add_string)
	for kkk:=range add_string{
		
	final_string = append(final_string, add_string[kkk])
	}





	}

  for i := 0; i < lll.mapers_number; i++ {
		
		
		dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		lp := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(i))
		os.Remove(lp)
	
		for j := 0; j < lll.reducers_number; j++ {

			dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		mapFile := path.Join(dir1, "MAP_"+file1+"_"+strconv.Itoa(i))
		dir2 := path.Dir(mapFile)
		file2 := path.Base(mapFile)
			
			os.Remove(path.Join(dir2, file2+"_"+strconv.Itoa(j)))
		}
	}


        file,_ := os.Create(lll.file_result)
	
	for kkkk:= range s{
	fmt.Fprint(file, final_string[kkkk]+" "+strconv.Itoa(s[kkkk])+"\n")

	}
		file.Close()


	
      
	for i := 0; i < lll.reducers_number; i++ {
		dir1 := path.Dir(lll.file)
		file1 := path.Base(lll.file)
		p:=path.Join(dir1, "RED_"+file1+"_"+strconv.Itoa(i))
	
		
		os.Remove(p)
	}
 
	
	}










func WordCount_GO(inFile string, outFile string) {

dt,_ := ioutil.ReadFile(inFile)


  ss := strings.Fields(string(dt))
   


    ts := make(map[string]int)

    for _, wd := range ss {
        _, ppp := ts[wd]

        if ppp == true {
          ts[wd] += 1
        } else {
          ts[wd] = 1
        } 
    } 




p := make(Key_ValueList, len(ts))

	i := 0
	for k, v := range ts {
		p[i] = Key_Value{k, v}
		i++
	
	}

	sort.Sort(p)


var s []int
var final_string []string
for _, v := range p {
s = append(s, v.Value)
}


result := removeDuplicates(s)
    

for kk:=range result{

var add_string []string
for _, v := range p {
    
    if v.Value == result[kk] { 
     
	add_string = append(add_string, v.Key)
	
      
    }
  }
sort.Strings(add_string)
for kkk:=range add_string{
	
final_string = append(final_string, add_string[kkk])
}





}

file,_ := os.Create(outFile)

     
for kkkk:= range s{
fmt.Fprint(file, final_string[kkkk]+" "+strconv.Itoa(s[kkkk])+"\n")

}


}



func WordCount_UNIX(inFile string, outFile string) {

	
	cmd:="cat" +" "+ inFile + "| tr '[:space:]' '[\\n*]' | grep -v \"^\\s*$\" | sort | uniq -c | LC_COLLATE=C sort -k1,1bnr -k2,2 | awk '{print $2\" \"$1}' >"+ outFile

        exec.Command("bash","-c",cmd).Output()

}

func main() {
	if len(os.Args) < 4 || len(os.Args) > 6{
		fmt.Printf("%s:\n\tUsage: bin/wordcount run-mode inFile outFile <mapers_number> <reducers_number>\n", os.Args[0])
		fmt.Printf("\tRun Modes:\n\t\t1 Unix Pipeline\n\t\t2 Simple wordcount in GO\n\t\t3 MapReduce Sequential\n\t\t4 MapReduce SMP\n\t\t5 MapReduce DMP\n")
		return
	}

	switch os.Args[1] {
  case "1":
    WordCount_UNIX(os.Args[2], os.Args[3])
  case "2":
    WordCount_GO(os.Args[2], os.Args[3])
  case "3":
    mapers_number, _ := strconv.Atoi(os.Args[4])
    reducers_number, _ := strconv.Atoi(os.Args[5])
    WordCount_MR_S(os.Args[2], os.Args[3], mapers_number, reducers_number)
  case "4":
    mapers_number, _ := strconv.Atoi(os.Args[4])
    reducers_number, _ := strconv.Atoi(os.Args[5])
    WordCount_MR_SMP(os.Args[2], os.Args[3], mapers_number, reducers_number)
  case "5":
    mapers_number, _ := strconv.Atoi(os.Args[4])
    reducers_number, _ := strconv.Atoi(os.Args[5])
    WordCount_MR_DMP(os.Args[2], os.Args[3], mapers_number, reducers_number)
  default:
    fmt.Printf("Unknown run-mode\n")	
  }
}
