import { TestBed } from '@angular/core/testing';

import { GongxsdspecificService } from './gongxsdspecific.service';

describe('GongxsdspecificService', () => {
  let service: GongxsdspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongxsdspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
