import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from 'src/app/services/local-storage.service';
@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  public senderNameSelected: string;
  public senderList: any;
  public msgInput: string;
  public msgList: any;
  constructor(
    public http: HttpClient,
    private reqProto: ReqProto,
    private lSData: LocalStorageService
  ) { }

  ngOnInit() {
    this.lSData.set("isLogin", "false");
    // 获取所有发送私信的人
    this.getSanderList();
  }
  getSanderList() {
    this.reqProto.action = "POST";
    this.reqProto.data = {
      uid: 1,
    }
    this.http.post("/api/getSenderList", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.senderList = res.data;
      console.log(res);

    })
  }
  // 获取与某人的所有聊天记录
  getRecords(uid: number, username: string) {
    this.senderNameSelected = username;
    this.reqProto.data = {
      sander: uid,
      receiver: 1
    }
    this.http.post("/api/getRecords", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.msgList = res.data;
      console.log(res);
    })
  }
  // 回复私信
  send() {
    this.http.post("/api/sendMsg", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.msgList.push(res.data);
      console.log(res);
    })
  }
}
