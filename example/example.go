package example

type Example struct {
	value0 int    // ignored
	value1 int    `genprop:"get"`
	value2 int    `genprop:"get,set"`
	id     int    `genprop:"get"`
	api    string `genprop:"get"`
	url    string `genprop:"get"`
	http   string `genprop:"get"`
}
