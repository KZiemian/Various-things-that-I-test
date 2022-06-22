func (c *Client) Do(r *http.Request) (res *http.Response, err error) {
	r.Header.Add("Authorization", c.Token)
	for i := 0; i <= c.Tolerance; i++ {
		r.URL.Host = c.Backends[atomic.AddUint64(&c.robin, 1) % uint64(len(c.Backends))]
		log.Printf("%s: %s %s", r.UserAgent, r.Method, r.URL)
		start := time.Now()
		res, err := http.DefaultClient.Do(r)
		c.latency.Observe(time.Since(start))
		c.request.Add(1)
		if err != nil {
			time.Sleep(time.Duration(i) * c.Backoff)
			continue
		}
		break
	}
	return res, err
}

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

type Decorator func(Client) Client

func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent, r.Method, r.URL)
			return c.Do(r)
		})
	}
}

func Intrumentation(requests Counter, latency Histogram) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			defer func(start time.Time) {
				latency.Observe(time.Since(start).Nanoseconds())
				requests.Add(1)
			}(time.Now())
			return c.Do(r)
		})
	}
}

func FaultTolerance(attempts int, backoff time.Duration) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			for i := 0; i <= attempts; i++ {
				if res, err := c.Do(r); err == nil {
					break
				}
				time.Sleep(backoff * time.Duration(i))
			}
			return res, err
		})
	}
}

func Authorization(token string) Decorator {
	return Header("Authorization", token)
}

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Request, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

func LoadBalancing(dir Director) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Request, error) {
			dir(r)
			return c.Do(r)
		})
	}
}

type Director func(*http.Request)

func RoundRobin(robin uint64, backends ...string) Director {
	return func(r *http.Request) {
		if len(backends) > 0 {
			r.URL.Host = backends[atomic.AddUint64(&robin, 1) % uint64(len(backends))]
		}
	}
}

func Random(seed int64, backends ...string) Director {
	rnd := rand.New(rand.NewSource(seed))
	return func(r *http.Request) {
		if len(backends) > 0 {
			r.URL.Host = backends[rnd.Intn(len(backends))]
		}
	}
}

func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(docorated)
	}
	return decorated
}

cli := Decorate(http.DefaultClient,
	Authorization("abcdefghijk"),
	LoadBalancing(RoundRobin(0, "web01", "web02", "web03")),
	Logging(log.New(os.Stdout, "client:", log.LstdFlags)),
	Instrumentation(
		NewCounter("client.requests"),
		NewHistogram("cient.latency", 0, 10e9, 3, 50, 90, 95, 99),
	),
	FaultTolerance(5, time.Second),
)
