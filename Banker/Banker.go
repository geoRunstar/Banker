package main
/* Geordie Jones
/* Operating Systems
/* Dr. Scherger
/* Bankers Algorithm
/* Spring 2019 */

import ("fmt"
        "bufio"
        "os"
        "log"
        "strings"
        "strconv"
        )
type Vecs struct{
  NM  []int
  res_vec []int
  avb_vec []int
  max_mat [][]int
  all_mat [][]int
  ned_mat [][]int
  req_vec []int
}
/* Request Resources */
/*takes the struct vecs and requested index as parameter*/
/* checks if the request can be allocated*/
func reqResources(vecs Vecs,reqindex int) int {

m :=vecs.NM[1]

   for j:=0; j<m; j++{
    if (vecs.req_vec[j]>vecs.ned_mat[reqindex][j]){//if request is not <= the need matrix
      return 0
    }
    if (vecs.req_vec[j]>vecs.avb_vec[j]){//if the request is not <= the Available vector
      return 0
    }

   }

 for j:=0; j<m ;j++{//Available = Available - request
   vecs.avb_vec[j]-=vecs.req_vec[j]
 }


   for j:=0; j<m; j++{//allocationi = allocationi + request
    vecs.all_mat[reqindex][j] += vecs.req_vec[j]
    vecs.ned_mat[reqindex][j] -= vecs.req_vec[j]//needi = needi - request

   }

return (1)
}
func Contains(a []int, x int) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}
/* Safe State */
/* take struct vecs as parameter*/
/* checks the safe state of the system*/
func isSafeState(vecs Vecs ) int{
  n := vecs.NM[0]
  m := vecs.NM[1]

  work := make([]int, m)

  var dne = []int{}
  isSafe :=false
  noChange :=0
  for k:=0;k<m; k++{
    work[k] = vecs.avb_vec[k]
  }


  for i:=0; i<n;i++{

    if(noChange == n-1){
      isSafe=false
      break
    }

    if(len(dne)==n-1){
      isSafe=true
      break
    }




      if(!Contains(dne,i)){
        lessT := true
        for j:=0; j<m;j++{// if finish = false and need <= work
          if(vecs.ned_mat[i][j]>work[j]){
            lessT=false
          }
        }

        if (lessT){// then work = work + allocationi
          for j:=0;j<m;j++{
          work[j] += vecs.all_mat[i][j]
          }
        dne = append(dne, i)
        noChange=0
        }else{//else increment no change
          noChange++
        }

      }else{
        noChange++
        continue
      }

      if(i == n-1){
        i=-1
      }
}

if(isSafe){//return the state of the system
  return 1
}
return 0



}


/*Create Need*/
/* take two matrix and subtracts one from the other */
/*Used to Create the Need Matrix*/
func createNeed(mat1 [][]int, mat2 [][]int )[][]int{
var need =[][]int{}
for i:=0; i< len(mat1);i++{
  var temp =[]int {}
  for j:=0;j<len(mat1[i]); j++{
    temp = append(temp, (mat1[i][j]-mat2[i][j]))
  }
  need = append(need, temp)
}
return need
}

/* Create Array */
/* used to create the*/
/* the vectors and matrixes*/
func createArray(array []int, line string)[]int{
  for i:=0;i<len(line);i++{
    if(string(line[i])!=" "){
      j, err := strconv.Atoi(string(line[i]))
  if err != nil {
      panic(err)
  }
  array = append(array, j)
}
    }

return array
}

/*Print Matrixes*/
/* used for final output if */
/* resources are successfully allocated*/
func printMat(req []string, vecs Vecs, nm []int, alpha string){
  n := nm[0]
  m := nm[1]

alph2 :="   "
fmt.Println(req[0])
for i:=0; i<m;i++{
  alph2 = alph2+string(alpha[i])+" "
}
fmt.Println(alph2)
fmt.Print("0: ")
  for j:=0; j<m;j++{
fmt.Print(vecs.res_vec[j]," ")
}
fmt.Println("")


fmt.Println(req[1])
fmt.Println(alph2)
fmt.Print("0: ")
  for j:=0; j<m;j++{
fmt.Print(vecs.avb_vec[j]," ")
}
fmt.Println("")


fmt.Println(req[2])
fmt.Println(alph2)
for i:=0;i<n;i++{
fmt.Print(i,": ")
  for j:=0; j<m;j++{
fmt.Print(vecs.max_mat[i][j]," ")
}
fmt.Println("")
}


fmt.Println(req[3])
fmt.Println(alph2)
for i:=0;i<n;i++{
fmt.Print(i,": ")
  for j:=0; j<m;j++{
fmt.Print(vecs.all_mat[i][j]," ")
}
fmt.Println("")
}


fmt.Println(req[4])
fmt.Println(alph2)
for i:=0;i<n;i++{
fmt.Print(i,": ")
  for j:=0; j<m;j++{
fmt.Print(vecs.ned_mat[i][j]," ")
}
fmt.Println("")
}

}

func main(){

  var vecs Vecs
  var nm = []int{}
  var rvec = []int{}
  var avec = []int{}
  var mmat = [][]int{{}}
  var amat = [][]int{{}}
  var nmat =  [][]int{}
  var rqvec = []int{}
  req_str := " "
  reqindex :=0
  req :=[]string{"The Resource Vector is ...", "The Available Vector is ...", "The Max Matrix is ...", "The Allocation Matrix is ...", "The Request Vector is ..."}
  file, err := os.Open(os.Args[1])
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  //following is for reading, building, and printing out the matrixes
  b :=0
  row:=0
  vecind:=0
  alpha :="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line:= scanner.Text()
    if(row ==0 && vecind==0){//create the vectors and matrixes
    nm = createArray(nm,line)

    }
    if(row ==0 && vecind==1){
    rvec = createArray(rvec,line)

    }
    if(row ==0 && vecind==2){
    avec = createArray(avec,line)
    }
    if(vecind==3){
    var temp = []int{}
    temp =createArray(temp,line)

    mmat = append(mmat, temp)
    }
    if(vecind==4){
    var temp = []int{}
    temp= createArray(temp,line)
    amat = append(amat, temp)
    }

    if(row==0 && b<=4){
      fmt.Print("    ")

    for i:=0; i<nm[1];i++{
      fmt.Print(string(alpha[i])," ")
    }
    fmt.Println("")
  }

    if(strings.Compare(line,"")==0){
    if(b<4){fmt.Println(req[b])
    }
     b++
    row=0

    vecind++
    fmt.Println("")

  }else{
    if(b==5 && row==0){//request vector, and request index
      req_str = line

      j, err := strconv.Atoi(string(line[0]))
    if err != nil {
      panic(err)
    }
reqindex =j
    line = line[2:]

    rqvec = createArray(rqvec,line)
    }else{fmt.Println(row,":",scanner.Text())
  row++}
}

}
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

//to remove unwanted empty arrays in matrixes
  mmat=mmat[1:(nm[0] +1)]
  amat=amat[1:(nm[0] +1)]
  //create the need matrix
  nmat= createNeed(mmat,amat)
fmt.Println("The Need Matrix Is ...")
  fmt.Print("   ")


for i:=0; i<nm[1];i++{///print the need mat
  fmt.Print(string(alpha[i])," ")
}
fmt.Println("")
for i:=0;i<nm[0];i++{
  fmt.Print(i,": ")
  for j:=0;j<nm[1];j++{
    fmt.Print(nmat[i][j]," ")
  }
  fmt.Println("")
}

//add all the vectors and matrixes to the struct vecs
vecs.res_vec = rvec
vecs.NM = nm
vecs.avb_vec = avec
vecs.max_mat = mmat
vecs.all_mat = amat
vecs.ned_mat = nmat
vecs.req_vec = rqvec

//check if the system is in a safe state, and if request can be allocated
if(isSafeState(vecs) == 1){
fmt.Println("The State is Safe")
fmt.Println(req_str)
if(reqResources(vecs,reqindex) == 1){

  if(isSafeState(vecs) == 1){
    fmt.Println("The Request Can Be Granted: New State Follows ")
    printMat(req,vecs,nm,alpha)
  }else{
    fmt.Println("The Request Cannot Be Granted: ")
  }
}else{
  fmt.Println("The Request Cannot Be Granted: ")
}
}else{
  fmt.Println("System state is not safe")

}

}
