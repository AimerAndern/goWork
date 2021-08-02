package leetcode
func isThree(n int) bool {
	j :=0
	for i :=1;i<=n;i+=1{
		if n%i==0{
			j+=1
		}
	}
	return j ==3
}