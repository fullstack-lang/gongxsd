import { TestBed } from '@angular/core/testing';

import { ReqifspecificService } from './reqifspecific.service';

describe('ReqifspecificService', () => {
  let service: ReqifspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ReqifspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
