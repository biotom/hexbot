package service_test



//test http methods hitting github hexbot API
//
//
//ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	_, err  := fmt.Fprintln(w, "#228B22")
//	if err != nil {
//		t.Fatal(err)
//	}
//}))
//defer ts.Close()
//
//
////this happens inside the colour service
//resp, err := http.Get(ts.URL)
//if err != nil {
//t.Fatal(err)
//}
//
//resp, err :=
//
//
//colour, err := ioutil.ReadAll(resp.Body)
//resp.Body.Close()
//if err != nil{
//t.Fatal(err)
//}
//
