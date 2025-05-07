import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongxsdSpecificComponent } from './gongxsd-specific.component';

describe('GongxsdSpecificComponent', () => {
  let component: GongxsdSpecificComponent;
  let fixture: ComponentFixture<GongxsdSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongxsdSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongxsdSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
