import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserInfoServiceService {

  private titleSource = new Subject
  //获得一个Observable
  titleObservable =this.titleSource.asObservable();
  constructor() { }
  //发射数据，当调用这个方法的时候，Subject就会发射这个数据，所有订阅了这个Subject的Subscription都会接受到结果
  emitTitle(title: string) {
      this.titleSource.next(title);
  }

}
