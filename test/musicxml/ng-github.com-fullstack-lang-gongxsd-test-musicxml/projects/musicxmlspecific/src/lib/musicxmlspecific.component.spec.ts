import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MusicxmlspecificComponent } from './musicxmlspecific.component';

describe('MusicxmlspecificComponent', () => {
  let component: MusicxmlspecificComponent;
  let fixture: ComponentFixture<MusicxmlspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MusicxmlspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(MusicxmlspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
