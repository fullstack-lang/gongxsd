import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BooksspecificComponent } from './booksspecific.component';

describe('BooksspecificComponent', () => {
  let component: BooksspecificComponent;
  let fixture: ComponentFixture<BooksspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BooksspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(BooksspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
