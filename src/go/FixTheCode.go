type Product struct {
	Count int
}

func (o *Product) Increase(n int) {
	o.Count += n
}

func (o *Product) Decrease(n int) {
	o.Count -= n
}

func (o *Product) String() string {
	return fmt.Sprintf("%v", o.Count)
}

func main() {
	var product Product
	product.Increase(10)
	product.Decrease(5)
	fmt.Println(product)
}

//INSTRUCTIONS: FIND OUT WHATS WRONG WITH THE ABOVE CODE AND FIX IT
