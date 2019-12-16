import { TestBed } from '@angular/core/testing';

import { UserInfoServiceService } from './user-info-service.service';

describe('UserInfoServiceService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: UserInfoServiceService = TestBed.get(UserInfoServiceService);
    expect(service).toBeTruthy();
  });
});
