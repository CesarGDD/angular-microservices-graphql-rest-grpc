import { Injectable } from "@angular/core";
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from "@angular/router";
import { Observable } from "rxjs";
import { AuthService } from "./auth.service";

@Injectable({providedIn: 'root'})
export class AuthGuard implements CanActivate {
    constructor (
        private authSvc: AuthService,
        private router: Router
    ) {}
    canActivate(
        route: ActivatedRouteSnapshot, 
        router: RouterStateSnapshot)
        : boolean | UrlTree | Promise<boolean | UrlTree> | Observable<boolean | UrlTree> {
            if (this.authSvc.token) {
                return true
            }
            else {
                return this.router.createUrlTree(['/auth'])
            }
        }
}