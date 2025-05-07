import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongxsdspecificComponent } from './gongxsdspecific.component';

describe('GongxsdspecificComponent', () => {
  let component: GongxsdspecificComponent;
  let fixture: ComponentFixture<GongxsdspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongxsdspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongxsdspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
