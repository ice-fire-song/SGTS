import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from '../../services/local-storage.service';
import { UserInfoServiceService } from 'src/app/services/user-info-service.service';
import { Subscription } from 'rxjs';
@Component({
  selector: 'app-good-detail',
  templateUrl: './good-detail.component.html',
  styleUrls: ['./good-detail.component.css']
})
export class GoodDetailComponent implements OnInit {
  public good: any;
  public imgList: any=[{save_path:"assets/images/奥地利.jpg"},{save_path:"assets/images/奥地利.png"},{save_path:"assets/images/奥地利.png"}];
  public gid: number;
  public isFavour: boolean = false;
  public favourStatusImg: string = "assets/images/取消收藏.png"
  private userInfo: any;
  subscription: Subscription;

  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    public http: HttpClient,
    private reqProto: ReqProto,
    private lSData: LocalStorageService,
    private uIService: UserInfoServiceService
  ) {
    this.subscription = this.uIService.titleObservable.subscribe((src:any) => {
      console.log('得到的title:' + src);
      this.userInfo = JSON.parse(src);
      console.log("userinfo:",this.userInfo)
      if (this.userInfo != null) {
        this.good = this.lSData.getObject("goodInfo");
        this.getGoodImg(); 
        this.reqProto.action = 'POST';
        this.reqProto.data = {
          gid: this.good.gid,
          uid: this.userInfo.uid
        }
        this.http.post("/api/seeFavourStatus", this.reqProto, this.httpOptions).subscribe((res: any) => {
          console.log("/api/seeFavourStatus", res);
          if (res.data.isFavour) {
            this.isFavour = true;
            this.favourStatusImg = "assets/images/收藏.png";
          }
        })
      }
    });
   }

  ngOnInit() {
  }
  // 添加收藏
  addFavour() {
    if (!this.isFavour) {
      this.reqProto.action = "POST"
      this.reqProto.data = {
        uid: this.userInfo.uid,
        gid: this.good.gid
      }
      this.http.post("/api/addFavour", this.reqProto, this.httpOptions).subscribe((res: any) => {
        console.log(res);
        if (res.isSuccess) {
          this.isFavour = true;
          this.favourStatusImg = "assets/images/收藏.png"
        }
      })
    } else {
      this.isFavour = false;
      this.favourStatusImg = "assets/images/取消收藏.png"
    }
  }
 
  // 获取货品所有图片
  getGoodImg() {
    this.reqProto.action = "POST";
    this.reqProto.data = {
      gid: this.gid
    }
    this.http.post("/api/getGoodImg", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.imgList = res.data;
      console.log(res);
    })
  }
  // 发私信
  public message_content: string = "";
  sendLetter() {
    this.reqProto.action = 'POST';
    this.reqProto.data = {
      user_id: this.userInfo.uid,
      friend_id: this.good.uid,
      message_content: this.message_content
    }
    console.log("this.resProto", this.reqProto)
    this.http.post("/api/sendPrivateLetter", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log("sendPrivateLetter", res)
    })
  }
  // 获取时间
  cutTime(time: string): string {
    if (time === undefined || time == null) {
      console.warn("cutTime()获取time为undefined或null")
    }
    return time.substr(0, 10); //+ ' ' + time.substr(11, 5);
  }
}
