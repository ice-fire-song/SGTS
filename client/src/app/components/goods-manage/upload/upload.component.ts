import { Component, OnInit,ViewChild,ElementRef } from '@angular/core';
import { NzMessageService  } from 'ng-zorro-antd';
import {  UploadFile } from 'ng-zorro-antd';

import { Router } from '@angular/router'
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto, Uploadfile } from 'src/app/services/proto'
import { LocalStorageService } from 'src/app/services/local-storage.service';
import { UserInfoServiceService } from 'src/app/services/user-info-service.service';
import { Subscription } from 'rxjs';
import axios from 'axios';//引入第三方模块进行数据请求

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})

export class UploadComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  public uid: number;
  public gname: string;          
  public gprice: number;   
  public gdetail: string;           
  public category_id: number = 0;        
  public gt_id: number;   
  private userInfo: any;
  subscription: Subscription;
  public priceHidden: boolean = false;
  constructor(
    private nzMessageService: NzMessageService,
    public http: HttpClient,
    private reqProto: ReqProto,
    private router: Router,
    private lSData: LocalStorageService,
    private uIService: UserInfoServiceService,
  ) {

    this.subscription = this.uIService.titleObservable.subscribe((src:any) => {
      console.log('得到的userInfo:' + src);
      this.userInfo = JSON.parse(src);
      console.log("userinfo:",this.userInfo)
      if (this.userInfo != null) {
        this.uid = this.userInfo.uid;
      } else {
        console.log("upload: 未接收到来自navigation的userInfo，则用户不处于登录状态或出错，将跳转到login"); 
        this.router.navigate(["/login"]);
      }
    });
 
   }

  ngOnInit() {
  }
  // 注意，所有图片均存入“知晓云”上，getImgURL()内主要是图片存储的流程
  public img_url: string = "";
  getImgURL(file: any) {
    // 第一步 获取上传文件所需授权凭证和上传地址
    let httpOptions1 = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' , 'Authorization': "Bearer 5dab952fe95b16a0ff7e2f20cb27183233ad9c0b"})
    };
    let data1 = {
        filename: file.name,
        category_id:"5de8a4e8a87859735bffea80"
    }
    let policy: string="";
    let authorization: string = "";
    let upload_url: string = "";
    console.log("first data:",data1)
     // 访问的地址：https://cloud.minapp.com/oserve/v2.1/upload/
    this.http.post("/first/oserve/v2.1/upload/", data1, httpOptions1).subscribe((res: any) => {
      console.log("firstHttp", res)
      policy = res.policy;
      authorization = res.authorization;
      upload_url = res.upload_url;

      // 第二步 上传图片，获取图片的url
      // 当前 upload_url : https://v0.api.upyun.com/cloud-minapp-30130
      const formData = new FormData();
      formData.append('authorization', authorization)
      formData.append('policy', policy)
      formData.append('file', file);
      axios({
        method: 'post',
        url: '/second/cloud-minapp-30130',
        data: formData,
        headers: {
          'Content-Type': 'multipart/form-data',
          'Authorization': "Bearer 5dab952fe95b16a0ff7e2f20cb27183233ad9c0b"
        },
        
      })
        .then(
          res => {
            console.log("res::", res)
            this.img_url = res.data.url
            if (!this.isFirstImg) {
              this.isFirstImg = true;
              this.first_img_path = this.img_base_path + this.img_url;
            }
            let uploadfile: Uploadfile = {};
            uploadfile.save_path = this.img_base_path + this.img_url;
            uploadfile.image_name = file.name;
            uploadfile.image_size = res.data.file_size;
            this.imgList.push(uploadfile);
            console.log("imglist",this.imgList)
          })
        .catch(err => {
          console.log("error:", err)
        });
    });
    this.fileList = this.fileList.concat(file);
      this.fileList.forEach((file: any) => {
      });  return false; 
  }


  // 第三步
  // 每张图片访问的域名 : https://cloud-minapp-30130.cloud.ifanrusercontent.com/
  public img_base_path = "https://cloud-minapp-30130.cloud.ifanrusercontent.com/"
  fileList: UploadFile[] = [];
  imgList: Uploadfile[] = [];
  isFirstImg: boolean = false;
  beforeUpload = (file: UploadFile): boolean => {//图片预处理
    // 第一步
    this.getImgURL(file);
    console.log(this.imgList)
    return false;
  };
  public att_exchange =0//判断附件上传情况
  uploadsure1(): void {//上传时
    this.reqProto.action = 'POST';
    this.reqProto.data = this.imgList;
    this.http.post("/api/uploadImg", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log("uploadimg:",res)
    })

  }
  // 货品类别（免费商品，商品，需求）tab切换
  switch(category_id: number) {
    if (category_id == 1) {
      this.priceHidden = true;
    } else {
      this.priceHidden = false;
    }
  }
  // 类型
  selectedProvince = '其他';
  provinceData = ['图书', '文具','生活用品', '电子产品','化妆用品', '服装鞋包','其他'];
  provinceChange(value: string): void { // 选择不用的货品类型
    for (let i in this.provinceData) {
      if (this.provinceData[i] == value) {
        this.gt_id = Number(i) + 1;
      }
    }
    console.log(value)
  }
  // 标签
  tags = [];
  inputVisible = false;
  inputValue = '';
  @ViewChild('inputElement', { static: false }) inputElement: ElementRef;

  handleClose(removedTag: {}): void {
    this.tags = this.tags.filter(tag => tag !== removedTag);
  }
  sliceTagName(tag: string): string {
    const isLongTag = tag.length > 20;
    return isLongTag ? `${tag.slice(0, 20)}...` : tag;
  }

  showInput(): void {
    this.inputVisible = true;
    setTimeout(() => {
      this.inputElement.nativeElement.focus();
    }, 10);
  }

  handleInputConfirm(): void {
    if (this.inputValue && this.tags.indexOf(this.inputValue) === -1) {
      this.tags = [...this.tags, this.inputValue];
    }
    this.inputValue = '';
    this.inputVisible = false;
  }

  // 上传
  cancel(): void {
    this.nzMessageService.info('click cancel');
  }
  first_img_path = "https://cloud-minapp-30130.cloud.ifanrusercontent.com/1ieRyUkGa0KrZREk.png"
  confirm(): void {
    this.reqProto.action = 'POST';
    this.reqProto.data = {
      uid: this.uid,
      gname: this.gname,          
      gprice: this.gprice,  
      gdetail: this.gdetail,         
      category_id: this.category_id,       
      gt_id: this.gt_id,
      tabs: this.tags,
      images: this.imgList,
      first_img_path: this.first_img_path
    }
    console.log("data",this.reqProto)
    this.http.post("/api/uploadGood", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log(res);
      if (res.data== true) {
        this.nzMessageService.success("货品发布成功");
        // this.uploadsure1();
      } else {
        this.nzMessageService.error("货品发布失败，请重新发布");
      }
    })
  }
}



