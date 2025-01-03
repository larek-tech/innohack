/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as SignupImport } from './routes/signup'
import { Route as LoginImport } from './routes/login'
import { Route as DashImport } from './routes/dash'
import { Route as ChatImport } from './routes/chat'
import { Route as IndexImport } from './routes/index'

// Create/Update Routes

const SignupRoute = SignupImport.update({
  id: '/signup',
  path: '/signup',
  getParentRoute: () => rootRoute,
} as any)

const LoginRoute = LoginImport.update({
  id: '/login',
  path: '/login',
  getParentRoute: () => rootRoute,
} as any)

const DashRoute = DashImport.update({
  id: '/dash',
  path: '/dash',
  getParentRoute: () => rootRoute,
} as any)

const ChatRoute = ChatImport.update({
  id: '/chat',
  path: '/chat',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/chat': {
      id: '/chat'
      path: '/chat'
      fullPath: '/chat'
      preLoaderRoute: typeof ChatImport
      parentRoute: typeof rootRoute
    }
    '/dash': {
      id: '/dash'
      path: '/dash'
      fullPath: '/dash'
      preLoaderRoute: typeof DashImport
      parentRoute: typeof rootRoute
    }
    '/login': {
      id: '/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof LoginImport
      parentRoute: typeof rootRoute
    }
    '/signup': {
      id: '/signup'
      path: '/signup'
      fullPath: '/signup'
      preLoaderRoute: typeof SignupImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '/chat': typeof ChatRoute
  '/dash': typeof DashRoute
  '/login': typeof LoginRoute
  '/signup': typeof SignupRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '/chat': typeof ChatRoute
  '/dash': typeof DashRoute
  '/login': typeof LoginRoute
  '/signup': typeof SignupRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/chat': typeof ChatRoute
  '/dash': typeof DashRoute
  '/login': typeof LoginRoute
  '/signup': typeof SignupRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths: '/' | '/chat' | '/dash' | '/login' | '/signup'
  fileRoutesByTo: FileRoutesByTo
  to: '/' | '/chat' | '/dash' | '/login' | '/signup'
  id: '__root__' | '/' | '/chat' | '/dash' | '/login' | '/signup'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  ChatRoute: typeof ChatRoute
  DashRoute: typeof DashRoute
  LoginRoute: typeof LoginRoute
  SignupRoute: typeof SignupRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  ChatRoute: ChatRoute,
  DashRoute: DashRoute,
  LoginRoute: LoginRoute,
  SignupRoute: SignupRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/chat",
        "/dash",
        "/login",
        "/signup"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/chat": {
      "filePath": "chat.tsx"
    },
    "/dash": {
      "filePath": "dash.tsx"
    },
    "/login": {
      "filePath": "login.tsx"
    },
    "/signup": {
      "filePath": "signup.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
