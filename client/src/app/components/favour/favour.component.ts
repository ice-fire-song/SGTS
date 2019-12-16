import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from 'src/app/services/local-storage.service';
import { UserInfoServiceService } from 'src/app/services/user-info-service.service';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';
import { NzMenuService, NzMessageService } from 'ng-zorro-antd';
@Component({
  selector: 'app-favour',
  templateUrl: './favour.component.html',
  styleUrls: ['./favour.component.css']
})
export class FavourComponent implements OnInit {
  public dirList: any;
  public goodList: any;
  public foldername = "默认收藏夹";
  public inputValue = "";
  public markFdid: number;
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  subscription: Subscription;
  public userInfo: any;
  public uid: Number;
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

        this.reqProto.action = "POST";
        this.reqProto.data = {uid: this.uid}
        this.http.post("/api/getFolder", this.reqProto, this.httpOptions).subscribe((response: any) => {
          this.dirList = response.data;
          console.log(response)
          if (this.dirList == null) {
            console.log("获取收藏夹失败");
            return
          }
          this.markFdid = this.dirList[0].fd_id;
          this.getGoodsOfDir(this.dirList[0].fd_id);
        });
      } else {
        console.log("favour: 未接收到来自navigation的userInfo，则用户不处于登录状态或出错，将跳转到login"); 
        this.router.navigate(["/login"]);
      }
    });
   }

  ngOnInit() {
    this.lSData.set("isLogin", "false");
  }

  getGoodsOfDir(fdid: number) {
    const data = {
      fdid: fdid,
      uid: this.uid,
      key: this.inputValue
    }
    this.reqProto.data = data
    this.http.post("/api/getFavourGoods", this.reqProto, this.httpOptions).subscribe((response: any) => {
      this.goodList = response.data;
      console.log(response)
    })
  }
  switch(fdid:number, foldername:string) {
    this.foldername = foldername;
    console.log(fdid, " ", foldername);
    this.getGoodsOfDir(fdid);
    this.markFdid = fdid; 
  }
  search() {
    this.getGoodsOfDir(this.markFdid);
  }
  // 删除收藏夹
  deleteDir(i: number, fd_id: number) {
    console.log(i," ",fd_id)
    if (i == 0) {
      console.log("默认收藏夹不能删除");
      this.nzMessageService.warning("默认收藏夹不能删除");
      return
    }
    this.reqProto.action = "POST";
    this.reqProto.data = {
      fd_id: fd_id
    }
    console.log("reqProto", this.reqProto)
    this.http.post("/api/deleteDir", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log(res)
      if (res.data.isSuccess) {
        this.dirList.splice(0, this.dirList.length - 1);
        this.nzMessageService.info('删除成功');
      
      } else {
        this.nzMessageService.info('删除失败');
      }
      console.log(res);
    })
  }
  // 取消按钮
  cancel(): void {
    this.nzMessageService.info('取消成功');
  }
 
  confirm(fid: number): void {
    console.log("fid", fid)
 
    this.reqProto.action = "POST";
    this.reqProto.data = {
      fid: fid
    }
    this.http.post("/api/deleteFavour", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log(res)
      if (res.data.isSuccess) {
        this.goodList.splice(0, this.goodList.length - 1);
        this.nzMessageService.info('取消成功');
      
      } else {
        this.nzMessageService.info('取消失败');
      }
      console.log(res);
    })
  }
  // 获取时间
  cutTime(time: string): string {
    if (time === undefined || time == null) {
      console.warn("cutTime()获取time为undefined或null")
    }
    return time.substr(0, 10) + ' ' + time.substr(11, 5);
  }
}
