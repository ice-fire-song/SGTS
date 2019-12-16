import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { ReqProto } from 'src/app/services/proto';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd';
import { LocalStorageService } from 'src/app/services/local-storage.service';
import { UserInfoServiceService } from 'src/app/services/user-info-service.service';

@Component({
  selector: 'app-personal-center',
  templateUrl: './personal-center.component.html',
  styleUrls: ['./personal-center.component.css']
})
export class PersonalCenterComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  subscription: Subscription;
  public userInfo: any;
  public uid: Number;
  public username: string;
  public label: string;
  public mailbox: string;
  public head_sculpture_path: string;
  public oldPassword: string="";
  public newPassword: string="";
  public correctNewPassword: string="";

  constructor(
    public http: HttpClient,
    private reqProto: ReqProto,
    private router: Router,
    private nzMessageService: NzMessageService,
    private lSData: LocalStorageService,
    private uIService: UserInfoServiceService
  ) {
    this.subscription = this.uIService.titleObservable.subscribe((src:any) => {
      console.log('得到的userInfo:' + src);
      this.userInfo = JSON.parse(src);
      console.log("userinfo:",this.userInfo)
      if (this.userInfo != null) {
        this.uid = this.userInfo.uid;
        this.username = this.userInfo.username;
        this.label = this.userInfo.label
        this.mailbox = this.userInfo.mailbox
        this.head_sculpture_path = this.userInfo.head_sculpture_path
      } else {
        console.log("favour: 未接收到来自navigation的userInfo，则用户不处于登录状态或出错，将跳转到login"); 
        this.router.navigate(["/login"]);
      }
    });
   }

  ngOnInit() {
  }
  // 提交修改按钮
  cancel(): void {
    this.nzMessageService.info('取消修改');
  }

  confirm(): void {
    console.log(this.label," ",this.head_sculpture_path," ",this.mailbox," ",this.newPassword)
    if (this.userInfo.label == this.label && this.userInfo.mailbox == this.mailbox && this.userInfo.head_sculpture_path == this.head_sculpture_path && this.newPassword == "") {
      this.nzMessageService.info("没有任何改动");
      return
    }
    if ((this.newPassword != "" || this.correctNewPassword != "") && this.newPassword != this.correctNewPassword) {
      this.nzMessageService.error("新密码不一致，请重新确认");
      return
    }
    this.reqProto.action = "POST";
    this.reqProto.data = {
      uid: this.userInfo.uid,
      label: this.label,
      mailbox: this.mailbox,
      head_sculpture_path: this.head_sculpture_path,
      oldPassword: this.oldPassword,
      newPassword: this.newPassword
    }
    this.http.post("/api/modifyUserInfo", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log("res:", res);
      if (res.data.isSuccess) {
        this.nzMessageService.info('修改成功');
      } else {
        this.nzMessageService.error('修改失败，请重新修改');
      }
    })
  }
}
