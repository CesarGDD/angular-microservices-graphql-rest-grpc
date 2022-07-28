import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './navbar/navbar.component';
import { PostcardComponent } from './postcard/postcard.component';
import { NewpostComponent } from './newpost/newpost.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { ApolloModule, APOLLO_OPTIONS } from 'apollo-angular';
import {HttpLink} from 'apollo-angular/http';
import { InMemoryCache } from '@apollo/client/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { GraphQLModule } from './graphql/graphql.module';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    PostcardComponent,
    NewpostComponent,
    HomeComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ApolloModule,
    HttpClientModule,
    FormsModule,
    GraphQLModule
  ],
  providers: [
    // {
    //   provide: APOLLO_OPTIONS,
    //   useFactory: (httpLink: HttpLink) => {
    //     return {
    //       cache: new InMemoryCache(),
    //       link: httpLink.create({
    //         uri: 'http://localhost:8080/query',
    //       }),
    //     };
    //   },
    //   deps: [HttpLink],
    // },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
