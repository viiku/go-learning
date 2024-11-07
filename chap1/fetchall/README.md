A goroutine is a concurrent function execution.
A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. 
The function main runs in a goroutine and the go statement creates additional goroutines.

The main function creates a channel of strings using make. For each command-line argument,
the go statement in the first range loop starts a new goroutine that calls fetch asynchronously
to fetch the URL using http.Get. The io.Copy function reads the body of the response and
discards it by writing to the ioutil.Discard output stream. Copy returns the byte count,
along with any error that occurred. As each result arrives, fetch sends a summary line on the
channel ch. The second range loop in main receives and prints those lines.
When one goroutine attempts a send or receive on a channel, it blocks until another goroutine
attempts the corresponding receive or send operation, at which point the value is transfer red
and both goroutines proceed. In this example, each fetch sends a value (ch <- expression) on
the channel ch, and main receives all of them (<-ch). Having main do all the printing ensures
that output from each goroutine is processed as a unit, with no danger of interleaving if two
goroutines finish at the same time.