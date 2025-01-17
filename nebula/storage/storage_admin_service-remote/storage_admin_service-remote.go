// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
        "../../github.com/vesoft-inc/nebula-go/v2/nebula/storage"
)

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  AdminExecResp transLeader(TransLeaderReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp addPart(AddPartReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp addLearner(AddLearnerReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp removePart(RemovePartReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp memberChange(MemberChangeReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp waitingForCatchUpData(CatchUpDataReq req)")
  fmt.Fprintln(os.Stderr, "  CreateCPResp createCheckpoint(CreateCPRequest req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp dropCheckpoint(DropCPRequest req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp blockingWrites(BlockingSignRequest req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp rebuildTagIndex(RebuildIndexRequest req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp rebuildEdgeIndex(RebuildIndexRequest req)")
  fmt.Fprintln(os.Stderr, "  GetLeaderPartsResp getLeaderParts(GetLeaderReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp checkPeers(CheckPeersReq req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp addAdminTask(AddAdminTaskRequest req)")
  fmt.Fprintln(os.Stderr, "  AdminExecResp stopAdminTask(StopAdminTaskRequest req)")
  fmt.Fprintln(os.Stderr, "  ListClusterInfoResp listClusterInfo(ListClusterInfoReq req)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.Transport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewHTTPPostClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewSocket(thrift.SocketAddr(net.JoinHostPort(host, portStr)))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.ProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := storage.NewStorageAdminServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "transLeader":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "TransLeader requires 1 args")
      flag.Usage()
    }
    arg186 := flag.Arg(1)
    mbTrans187 := thrift.NewMemoryBufferLen(len(arg186))
    defer mbTrans187.Close()
    _, err188 := mbTrans187.WriteString(arg186)
    if err188 != nil {
      Usage()
      return
    }
    factory189 := thrift.NewSimpleJSONProtocolFactory()
    jsProt190 := factory189.GetProtocol(mbTrans187)
    argvalue0 := storage.NewTransLeaderReq()
    err191 := argvalue0.Read(jsProt190)
    if err191 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.TransLeader(value0))
    fmt.Print("\n")
    break
  case "addPart":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "AddPart requires 1 args")
      flag.Usage()
    }
    arg192 := flag.Arg(1)
    mbTrans193 := thrift.NewMemoryBufferLen(len(arg192))
    defer mbTrans193.Close()
    _, err194 := mbTrans193.WriteString(arg192)
    if err194 != nil {
      Usage()
      return
    }
    factory195 := thrift.NewSimpleJSONProtocolFactory()
    jsProt196 := factory195.GetProtocol(mbTrans193)
    argvalue0 := storage.NewAddPartReq()
    err197 := argvalue0.Read(jsProt196)
    if err197 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.AddPart(value0))
    fmt.Print("\n")
    break
  case "addLearner":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "AddLearner requires 1 args")
      flag.Usage()
    }
    arg198 := flag.Arg(1)
    mbTrans199 := thrift.NewMemoryBufferLen(len(arg198))
    defer mbTrans199.Close()
    _, err200 := mbTrans199.WriteString(arg198)
    if err200 != nil {
      Usage()
      return
    }
    factory201 := thrift.NewSimpleJSONProtocolFactory()
    jsProt202 := factory201.GetProtocol(mbTrans199)
    argvalue0 := storage.NewAddLearnerReq()
    err203 := argvalue0.Read(jsProt202)
    if err203 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.AddLearner(value0))
    fmt.Print("\n")
    break
  case "removePart":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemovePart requires 1 args")
      flag.Usage()
    }
    arg204 := flag.Arg(1)
    mbTrans205 := thrift.NewMemoryBufferLen(len(arg204))
    defer mbTrans205.Close()
    _, err206 := mbTrans205.WriteString(arg204)
    if err206 != nil {
      Usage()
      return
    }
    factory207 := thrift.NewSimpleJSONProtocolFactory()
    jsProt208 := factory207.GetProtocol(mbTrans205)
    argvalue0 := storage.NewRemovePartReq()
    err209 := argvalue0.Read(jsProt208)
    if err209 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RemovePart(value0))
    fmt.Print("\n")
    break
  case "memberChange":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "MemberChange requires 1 args")
      flag.Usage()
    }
    arg210 := flag.Arg(1)
    mbTrans211 := thrift.NewMemoryBufferLen(len(arg210))
    defer mbTrans211.Close()
    _, err212 := mbTrans211.WriteString(arg210)
    if err212 != nil {
      Usage()
      return
    }
    factory213 := thrift.NewSimpleJSONProtocolFactory()
    jsProt214 := factory213.GetProtocol(mbTrans211)
    argvalue0 := storage.NewMemberChangeReq()
    err215 := argvalue0.Read(jsProt214)
    if err215 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.MemberChange(value0))
    fmt.Print("\n")
    break
  case "waitingForCatchUpData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "WaitingForCatchUpData requires 1 args")
      flag.Usage()
    }
    arg216 := flag.Arg(1)
    mbTrans217 := thrift.NewMemoryBufferLen(len(arg216))
    defer mbTrans217.Close()
    _, err218 := mbTrans217.WriteString(arg216)
    if err218 != nil {
      Usage()
      return
    }
    factory219 := thrift.NewSimpleJSONProtocolFactory()
    jsProt220 := factory219.GetProtocol(mbTrans217)
    argvalue0 := storage.NewCatchUpDataReq()
    err221 := argvalue0.Read(jsProt220)
    if err221 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.WaitingForCatchUpData(value0))
    fmt.Print("\n")
    break
  case "createCheckpoint":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateCheckpoint requires 1 args")
      flag.Usage()
    }
    arg222 := flag.Arg(1)
    mbTrans223 := thrift.NewMemoryBufferLen(len(arg222))
    defer mbTrans223.Close()
    _, err224 := mbTrans223.WriteString(arg222)
    if err224 != nil {
      Usage()
      return
    }
    factory225 := thrift.NewSimpleJSONProtocolFactory()
    jsProt226 := factory225.GetProtocol(mbTrans223)
    argvalue0 := storage.NewCreateCPRequest()
    err227 := argvalue0.Read(jsProt226)
    if err227 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CreateCheckpoint(value0))
    fmt.Print("\n")
    break
  case "dropCheckpoint":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "DropCheckpoint requires 1 args")
      flag.Usage()
    }
    arg228 := flag.Arg(1)
    mbTrans229 := thrift.NewMemoryBufferLen(len(arg228))
    defer mbTrans229.Close()
    _, err230 := mbTrans229.WriteString(arg228)
    if err230 != nil {
      Usage()
      return
    }
    factory231 := thrift.NewSimpleJSONProtocolFactory()
    jsProt232 := factory231.GetProtocol(mbTrans229)
    argvalue0 := storage.NewDropCPRequest()
    err233 := argvalue0.Read(jsProt232)
    if err233 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.DropCheckpoint(value0))
    fmt.Print("\n")
    break
  case "blockingWrites":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "BlockingWrites requires 1 args")
      flag.Usage()
    }
    arg234 := flag.Arg(1)
    mbTrans235 := thrift.NewMemoryBufferLen(len(arg234))
    defer mbTrans235.Close()
    _, err236 := mbTrans235.WriteString(arg234)
    if err236 != nil {
      Usage()
      return
    }
    factory237 := thrift.NewSimpleJSONProtocolFactory()
    jsProt238 := factory237.GetProtocol(mbTrans235)
    argvalue0 := storage.NewBlockingSignRequest()
    err239 := argvalue0.Read(jsProt238)
    if err239 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.BlockingWrites(value0))
    fmt.Print("\n")
    break
  case "rebuildTagIndex":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RebuildTagIndex requires 1 args")
      flag.Usage()
    }
    arg240 := flag.Arg(1)
    mbTrans241 := thrift.NewMemoryBufferLen(len(arg240))
    defer mbTrans241.Close()
    _, err242 := mbTrans241.WriteString(arg240)
    if err242 != nil {
      Usage()
      return
    }
    factory243 := thrift.NewSimpleJSONProtocolFactory()
    jsProt244 := factory243.GetProtocol(mbTrans241)
    argvalue0 := storage.NewRebuildIndexRequest()
    err245 := argvalue0.Read(jsProt244)
    if err245 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RebuildTagIndex(value0))
    fmt.Print("\n")
    break
  case "rebuildEdgeIndex":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RebuildEdgeIndex requires 1 args")
      flag.Usage()
    }
    arg246 := flag.Arg(1)
    mbTrans247 := thrift.NewMemoryBufferLen(len(arg246))
    defer mbTrans247.Close()
    _, err248 := mbTrans247.WriteString(arg246)
    if err248 != nil {
      Usage()
      return
    }
    factory249 := thrift.NewSimpleJSONProtocolFactory()
    jsProt250 := factory249.GetProtocol(mbTrans247)
    argvalue0 := storage.NewRebuildIndexRequest()
    err251 := argvalue0.Read(jsProt250)
    if err251 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RebuildEdgeIndex(value0))
    fmt.Print("\n")
    break
  case "getLeaderParts":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLeaderParts requires 1 args")
      flag.Usage()
    }
    arg252 := flag.Arg(1)
    mbTrans253 := thrift.NewMemoryBufferLen(len(arg252))
    defer mbTrans253.Close()
    _, err254 := mbTrans253.WriteString(arg252)
    if err254 != nil {
      Usage()
      return
    }
    factory255 := thrift.NewSimpleJSONProtocolFactory()
    jsProt256 := factory255.GetProtocol(mbTrans253)
    argvalue0 := storage.NewGetLeaderReq()
    err257 := argvalue0.Read(jsProt256)
    if err257 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetLeaderParts(value0))
    fmt.Print("\n")
    break
  case "checkPeers":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CheckPeers requires 1 args")
      flag.Usage()
    }
    arg258 := flag.Arg(1)
    mbTrans259 := thrift.NewMemoryBufferLen(len(arg258))
    defer mbTrans259.Close()
    _, err260 := mbTrans259.WriteString(arg258)
    if err260 != nil {
      Usage()
      return
    }
    factory261 := thrift.NewSimpleJSONProtocolFactory()
    jsProt262 := factory261.GetProtocol(mbTrans259)
    argvalue0 := storage.NewCheckPeersReq()
    err263 := argvalue0.Read(jsProt262)
    if err263 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CheckPeers(value0))
    fmt.Print("\n")
    break
  case "addAdminTask":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "AddAdminTask requires 1 args")
      flag.Usage()
    }
    arg264 := flag.Arg(1)
    mbTrans265 := thrift.NewMemoryBufferLen(len(arg264))
    defer mbTrans265.Close()
    _, err266 := mbTrans265.WriteString(arg264)
    if err266 != nil {
      Usage()
      return
    }
    factory267 := thrift.NewSimpleJSONProtocolFactory()
    jsProt268 := factory267.GetProtocol(mbTrans265)
    argvalue0 := storage.NewAddAdminTaskRequest()
    err269 := argvalue0.Read(jsProt268)
    if err269 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.AddAdminTask(value0))
    fmt.Print("\n")
    break
  case "stopAdminTask":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "StopAdminTask requires 1 args")
      flag.Usage()
    }
    arg270 := flag.Arg(1)
    mbTrans271 := thrift.NewMemoryBufferLen(len(arg270))
    defer mbTrans271.Close()
    _, err272 := mbTrans271.WriteString(arg270)
    if err272 != nil {
      Usage()
      return
    }
    factory273 := thrift.NewSimpleJSONProtocolFactory()
    jsProt274 := factory273.GetProtocol(mbTrans271)
    argvalue0 := storage.NewStopAdminTaskRequest()
    err275 := argvalue0.Read(jsProt274)
    if err275 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.StopAdminTask(value0))
    fmt.Print("\n")
    break
  case "listClusterInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ListClusterInfo requires 1 args")
      flag.Usage()
    }
    arg276 := flag.Arg(1)
    mbTrans277 := thrift.NewMemoryBufferLen(len(arg276))
    defer mbTrans277.Close()
    _, err278 := mbTrans277.WriteString(arg276)
    if err278 != nil {
      Usage()
      return
    }
    factory279 := thrift.NewSimpleJSONProtocolFactory()
    jsProt280 := factory279.GetProtocol(mbTrans277)
    argvalue0 := storage.NewListClusterInfoReq()
    err281 := argvalue0.Read(jsProt280)
    if err281 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ListClusterInfo(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
