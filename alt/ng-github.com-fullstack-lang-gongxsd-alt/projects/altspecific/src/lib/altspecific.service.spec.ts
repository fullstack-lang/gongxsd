import { TestBed } from '@angular/core/testing';

import { AltspecificService } from './altspecific.service';

describe('AltspecificService', () => {
  let service: AltspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AltspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
