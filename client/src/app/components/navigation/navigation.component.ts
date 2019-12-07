import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Router, NavigationEnd } from '@angular/router';
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from '../../services/local-storage.service';
import axios from 'axios';//引入第三方模块进行数据请求
import { filter } from 'rxjs/operators';
@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.css']
})
export class NavigationComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  public uid: number;
  Ishide: boolean = false;
  public userStatus: any;
  public userInfo: any =  {
    username: "",
    head_sculpture_path: "assets/images/head_img.jpg"
  };
  constructor(
    private http: HttpClient,
    private router: Router,
    private lSData: LocalStorageService,
    private reqProto: ReqProto,
  ) { 
    // this.router.events.subscribe((event)=>{
    //   console.log("event:", event)
      
    // })
    this.router.events.pipe(filter((event) => event instanceof NavigationEnd))
    .subscribe((event:NavigationEnd) => {
      console.log("event.url:", event.url)
      if (event.url != "/login") {
        this.Ishide = true;
        let that = this
        this.http.get("/api/login").subscribe((res: any) => {
          console.log("login:res:",res);
          this.userStatus = res.data;
          if (res.data.status) {
            this.reqProto.data = {
              username: res.data.username
            }
            this.http.post("/api/getUserInfo", this.reqProto, this.httpOptions).subscribe((res: any) => {
              console.log(res);
              this.userInfo = res.data;
            })
          } else {
            this.userInfo = {
              username: "",
              head_sculpture_path: "assets/images/head_img.jpg"
            }
          }
        })
    
      } else {
        this.Ishide = false;
      }
  });
  }

  ngOnInit() {
    // setTimeout("console.log('navigation 等待1s，为了判断isLogin')",1000);
    // let isLogin = this.lSData.get("isLogin");
    // console.log("isLogin: ",isLogin)
    // if (isLogin) {
    //   this.Ishide = false;
    //   console.log("现在在login页面")
     
    //   return
    // }

  }
  clickHeadImg() {
    if (this.userStatus.status == false) {
      this.router.navigateByUrl("/login");
    } else {
      this.router.navigateByUrl("/personalCenter")
    }
  }
  logout() {
    this.lSData.remove("isLogin");
    let that = this
    var url = "/api/logout"
    that.http.get(url,this.httpOptions).subscribe(res => {
      console.log(res);
      that.router.navigate(["/login"])
    });
  }
  jump(navName: string) {
    console.log(navName)
    if (navName == "登录") {
      this.router.navigate(['/login']);
    }else if (navName == "个人中心") {
      this.router.navigate(['/peronalCenter'], { queryParams: {uid: this.uid} });
    }else  if (navName == "私信") {
      this.router.navigate(['/chat'], { queryParams:  {uid: this.uid}  });
    }else  if (navName == "收藏") {
      this.router.navigate(['/favour'], { queryParams:  {uid: this.uid} });
    }else if (navName == "发布中心" || navName == "上传") {
      this.router.navigate(['/goodsmanage'], { queryParams:  {uid: this.uid} });
    }
    
  }
}
