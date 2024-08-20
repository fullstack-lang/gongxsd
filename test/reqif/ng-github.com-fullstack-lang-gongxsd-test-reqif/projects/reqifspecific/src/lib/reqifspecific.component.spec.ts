import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReqifspecificComponent } from './reqifspecific.component';

describe('ReqifspecificComponent', () => {
  let component: ReqifspecificComponent;
  let fixture: ComponentFixture<ReqifspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReqifspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ReqifspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
