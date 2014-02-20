package webserver;

/**
 * A simple writer which writes into a string.
 * @author Rémy MATHIEU
 */
type SimpleStringWriter struct {
    Value string;
}

func (s *SimpleStringWriter) Write(p []byte) (int, error) {
    s.Value = string(p);
    return len(p), nil;
}
