import { TestBed } from '@angular/core/testing';

import { BooksspecificService } from './booksspecific.service';

describe('BooksspecificService', () => {
  let service: BooksspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BooksspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
