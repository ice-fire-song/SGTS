import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
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
  public uid: Number = 1;
  constructor(
    public http: HttpClient,
    private reqProto: ReqProto,
  ) { }

  ngOnInit() {
    this.reqProto.action = "GET";
    this.reqProto.data = {uid: this.uid}
    this.http.post("/api/getFolder", this.reqProto, this.httpOptions).subscribe((response: any) => {
      this.dirList = response.data;
      console.log(response)
      this.markFdid = this.dirList[0].fdid;
      this.getGoodsOfDir(this.dirList[0].fdid);
    });

  }

  getGoodsOfDir(fdid: number) {
    const data = {
      fdid: fdid,
      uid: this.uid,
      key: this.inputValue
    }
    this.reqProto.data = data
    this.http.post("api/getFavourGoods", this.reqProto, this.httpOptions).subscribe((response: any) => {
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
  
  // 获取时间
  cutTime(time: string): string {
    if (time === undefined || time == null) {
      console.warn("cutTime()获取time为undefined或null")
    }
    return time.substr(0, 10) + ' ' + time.substr(11, 5);
  }
}
