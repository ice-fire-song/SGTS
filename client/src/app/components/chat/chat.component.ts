import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from 'src/app/services/local-storage.service';
import { UserInfoServiceService } from 'src/app/services/user-info-service.service';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd';
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
  public message_content: string = "";
  public msgList: any;
  public uid: number;
  public friend_id: number;
  private userInfo: any;
  subscription: Subscription;
  constructor(
    public http: HttpClient,
    private reqProto: ReqProto,
    private lSData: LocalStorageService,
    private router: Router,
    private uIService: UserInfoServiceService,
    private nzMessageService: NzMessageService
  ) { 
    this.subscription = this.uIService.titleObservable.subscribe((src:any) => {
      console.log('得到的userInfo:' + src);
      this.userInfo = JSON.parse(src);
      console.log("userinfo:",this.userInfo)
      if (this.userInfo != null) {
        this.uid = this.userInfo.uid;

        // 获取所有发送私信的人
        this.reqProto.action = "POST";
        this.reqProto.data = {uid: this.uid}
        this.http.post("/api/getSenderList", this.reqProto, this.httpOptions).subscribe((response: any) => {
          this.senderList = response.data;
          console.log(response);
        });
      } else {
        console.log("favour: 未接收到来自navigation的userInfo，则用户不处于登录状态或出错，将跳转到login"); 
        this.router.navigate(["/login"]);
      }
    });
  }

  ngOnInit() {
  }
  // 获取与某人的所有聊天记录
  getRecords(friend_id: number, username: string) {
    this.friend_id = friend_id;
    this.senderNameSelected = username;
    this.reqProto.data = {
      user_id: this.uid,
      friend_id: friend_id
    }
    this.http.post("/api/getRecords", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.msgList = res.data;
      console.log(res);
    })
  }
  // 回复私信
  send() {
    if (this.message_content.length == 0) {
      this.nzMessageService.warning("回复的内容不能为空");
      return;
    }
    this.reqProto.action = "POST";
    this.reqProto.data = {
      user_id: this.uid,
      friend_id: this.friend_id,
      message_content: this.message_content
    }
    this.http.post("/api/sendPrivateLetter", this.reqProto, this.httpOptions).subscribe((res: any) => {
      if (res.data.isSuccess) {
        this.nzMessageService.info("回复成功");
        this.msgList.push(res.data);
      } else {
        this.nzMessageService.error("回复失败");
     }
    })
  }
}
