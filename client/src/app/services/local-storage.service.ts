import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LocalStorageService {


  // public localStorage:any;
 
  constructor() {
    // if (!localStorage) {
    //   throw new Error('Current browser does not support Local Storage');
    // }
    // localStorage = localStorage;
  }
 
  public set(key:string, value:string):void {
    localStorage[key] = value;
  }
 
  public get(key:string):string {
    return localStorage[key] || false;
  }
 
  public setArr(key:string, value:Array<any>):void{
    localStorage[key] = value;
  }
 
  public setObject(key:string, value:any):void {
    localStorage[key] = JSON.stringify(value);
  }
 
  public getObject(key:string):any {
    return JSON.parse(localStorage[key] || '{}');
  }
 
  public remove(key:string):any {
    localStorage.removeItem(key);
  }
  public removeAll():any{
 
    localStorage.clear();
  }

}
