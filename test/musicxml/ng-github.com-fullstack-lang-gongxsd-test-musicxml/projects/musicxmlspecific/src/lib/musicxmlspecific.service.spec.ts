import { TestBed } from '@angular/core/testing';

import { MusicxmlspecificService } from './musicxmlspecific.service';

describe('MusicxmlspecificService', () => {
  let service: MusicxmlspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MusicxmlspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
