// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { ContentAPI } from './content-api'
import { Content, CopyContentToContentAPI } from './content'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class ContentService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ContentServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private contentsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.contentsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/markdown/go/v1/contents';
  }

  /** GET contents from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<ContentAPI[]> {
    return this.getContents(Name, frontRepo)
  }
  getContents(Name: string, frontRepo: FrontRepo): Observable<ContentAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<ContentAPI[]>(this.contentsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<ContentAPI[]>('getContents', []))
      );
  }

  /** GET content by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<ContentAPI> {
    return this.getContent(id, Name, frontRepo)
  }
  getContent(id: number, Name: string, frontRepo: FrontRepo): Observable<ContentAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.contentsUrl}/${id}`;
    return this.http.get<ContentAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched content id=${id}`)),
      catchError(this.handleError<ContentAPI>(`getContent id=${id}`))
    );
  }

  // postFront copy content to a version with encoded pointers and post to the back
  postFront(content: Content, Name: string): Observable<ContentAPI> {
    let contentAPI = new ContentAPI
    CopyContentToContentAPI(content, contentAPI)
    const id = typeof contentAPI === 'number' ? contentAPI : contentAPI.ID
    const url = `${this.contentsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ContentAPI>(url, contentAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ContentAPI>('postContent'))
    );
  }
  
  /** POST: add a new content to the server */
  post(contentdb: ContentAPI, Name: string, frontRepo: FrontRepo): Observable<ContentAPI> {
    return this.postContent(contentdb, Name, frontRepo)
  }
  postContent(contentdb: ContentAPI, Name: string, frontRepo: FrontRepo): Observable<ContentAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ContentAPI>(this.contentsUrl, contentdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted contentdb id=${contentdb.ID}`)
      }),
      catchError(this.handleError<ContentAPI>('postContent'))
    );
  }

  /** DELETE: delete the contentdb from the server */
  delete(contentdb: ContentAPI | number, Name: string): Observable<ContentAPI> {
    return this.deleteContent(contentdb, Name)
  }
  deleteContent(contentdb: ContentAPI | number, Name: string): Observable<ContentAPI> {
    const id = typeof contentdb === 'number' ? contentdb : contentdb.ID;
    const url = `${this.contentsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<ContentAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted contentdb id=${id}`)),
      catchError(this.handleError<ContentAPI>('deleteContent'))
    );
  }

  // updateFront copy content to a version with encoded pointers and update to the back
  updateFront(content: Content, Name: string): Observable<ContentAPI> {
    let contentAPI = new ContentAPI
    CopyContentToContentAPI(content, contentAPI)
    const id = typeof contentAPI === 'number' ? contentAPI : contentAPI.ID
    const url = `${this.contentsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<ContentAPI>(url, contentAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ContentAPI>('updateContent'))
    );
  }

  /** PUT: update the contentdb on the server */
  update(contentdb: ContentAPI, Name: string, frontRepo: FrontRepo): Observable<ContentAPI> {
    return this.updateContent(contentdb, Name, frontRepo)
  }
  updateContent(contentdb: ContentAPI, Name: string, frontRepo: FrontRepo): Observable<ContentAPI> {
    const id = typeof contentdb === 'number' ? contentdb : contentdb.ID;
    const url = `${this.contentsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<ContentAPI>(url, contentdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated contentdb id=${contentdb.ID}`)
      }),
      catchError(this.handleError<ContentAPI>('updateContent'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in ContentService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("ContentService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
