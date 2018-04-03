package main

// replicated storage with read and write quorums.

const (
	F           = 2
	N           = 5 // >= 2F + 1
	ReadQuorum  = F + 1
	WriteQuorum = N - F
)


type Server struct{}
type Request struct{}

func (s *Server) Write(req *Request) {

}

func (s *Server) Read(req *Request) *Data {

}
// Replicated write, returning after enough writes have succeeded.
func write(req *Request) {
	servers := []*Server{}
	done := make(chan bool, len(servers))

	for _, srv := range servers {
		go func(srv *Server) {
			srv.Write(req)
			done <- true
		}(srv)
	}

	for n := 0; n < WriteQuorum; n++ {
		<-done
	}
}

// Replicated read, returning after enough reads have been gathered.
func read(req *Request) {
	servers := []*Server{}
	replies := make(chan *Data, len(servers))

	for _, srv := range servers {
		go func(srv *Server) {
			replies <- srv.Read(key)
		}(srv)
	}

	var d *Data
	for n := 0; n < ReadQuorum; n++ {
		d = better(d, <-replies)
	}
}

for {
    select {
    case event := <-ui:
        // process user interface event
    case msg := <-server:
        // process server message
    case t := <-tick:
        // time has elapsed
    }
}