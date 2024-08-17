import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AltspecificComponent } from './altspecific.component';

describe('AltspecificComponent', () => {
  let component: AltspecificComponent;
  let fixture: ComponentFixture<AltspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AltspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AltspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
