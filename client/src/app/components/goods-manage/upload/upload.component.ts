import { Component, OnInit,ViewChild,ElementRef } from '@angular/core';
import { NzMessageService  } from 'ng-zorro-antd';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import {  UploadFile } from 'ng-zorro-antd';
import axios from 'axios'

import { Router } from '@angular/router'
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { ReqProto } from 'src/app/services/proto'
import { LocalStorageService } from 'src/app/services/local-storage.service';
@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent implements OnInit {
  public httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  public uid: number = 1;
  public gname: string;          
  public gprice: number;   
  public gdetail: string;           
  public category_id: number = 0;        
  public gt_id: number;   
  
  public priceHidden: boolean = false;
  constructor(
    private nzMessageService: NzMessageService,
    public http: HttpClient,
    private reqProto: ReqProto,
    private lSData: LocalStorageService
  ) { }

  ngOnInit() {
    this.lSData.set("isLogin", "false");
  }
  // tab切换
  switch(category_id: number) {
    if (category_id == 1) {
      this.priceHidden = true;
    } else {
      this.priceHidden = false;
    }
  }
  // 类别
  selectedProvince = '其他';
  provinceData = ['图书', '文具','生活用品', '电子产品','化妆用品', '服装鞋包','其他'];
  provinceChange(value: string): void {
    //this.selectedCity = this.cityData[value][0];
    for (let i in this.provinceData) {
      if (this.provinceData[i] == value) {
        this.gt_id = Number(i) + 1;
      }
    }
    console.log(value)
  }
  // 标签
  tags = ['Unremovable', 'Tag 2', 'Tag 3'];
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
  // 上传图片
  defaultFileList = [
    {
      uid: 1,
      name: 'xxx.png',
      status: 'done',
      url: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png',
      thumbUrl: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png'
    },
    {
      uid: 1,
      name: 'yyy.png',
      status: 'done',
      url: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png',
      thumbUrl: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png'
    }
  ];
  fileList2 = [...this.defaultFileList];
  // 上传
  cancel(): void {
    this.nzMessageService.info('click cancel');
  }

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
      images: this.fileList2
    }
    this.http.post("/api/uploadGood", this.reqProto, this.httpOptions).subscribe((res: any) => {
      console.log(res);
      if (res.data== true) {
        this.nzMessageService.success("货品发布成功");
      } else {
        this.nzMessageService.error("货品发布失败，请重新发布");
      }
    })
    this.nzMessageService.info('click confirm');
  }
}
