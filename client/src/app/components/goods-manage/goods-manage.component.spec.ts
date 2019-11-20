import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GoodsManageComponent } from './goods-manage.component';

describe('GoodsManageComponent', () => {
  let component: GoodsManageComponent;
  let fixture: ComponentFixture<GoodsManageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GoodsManageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GoodsManageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
