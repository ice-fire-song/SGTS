import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from '../../services/local-storage.service';
@Component({
  selector: 'app-good-detail',
  templateUrl: './good-detail.component.html',
  styleUrls: ['./good-detail.component.css']
})
export class GoodDetailComponent implements OnInit {
  public good: any;
  public imgList: any;
  public gid: number;
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor(
    public http: HttpClient,
    private reqProto: ReqProto,
    private lSData: LocalStorageService
  ) { }

  ngOnInit() {
    this.lSData.set("isLogin", "false");
    this.good = this.lSData.getObject("goodInfo");
    // this.getGoodInfo();
    // this.getGoodImg();
  }
  getGoodInfo() {
    this.reqProto.action = "POST";
    this.reqProto.data = {
      gid: this.gid
    }
    this.http.post("/api/getGoodInfo", this.reqProto, this.httpOptions).subscribe((res: any) => {
      this.good = res.data;
      console.log(res);
    })
  }
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
}
