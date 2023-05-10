import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import router from "@/router";
import {useUsersStore} from "@/stores/user";
import {useMenusStore} from "@/stores/menu";
import {assembledMenus, getUserAllActions} from "./api/menu";
import {useTabsStore} from "./stores/tab";
import {useActionsStore} from "@/stores/action";

NProgress.configure({ showSpinner: false })

const whiteList = ['/login']

router.beforeEach((to, from, next) => {
  NProgress.start()
  if(to.path === '/login'){
    next()
    NProgress.done()
  }else{
    const usersStore = useUsersStore()
    if (usersStore.token) {
      const menusStore = useMenusStore()
      if (menusStore.menus==null) {
        //assemble user menus and router
        Promise.all([assembledMenus(), getUserAllActions()]).then(function (results) {
          const actionStore = useActionsStore()
          const menus = results[0].data.result.items
          const userOwnedActions = results[1].data.result.actionName
          menusStore.setMenus(menus)

          menusStore.assembledRoutes()
          actionStore.setUserActions(userOwnedActions)
          next({ path: to.path })
        })

      }else{
        next()
      }
    }else{
      if (whiteList.indexOf(to.path) !== -1) {
        next()
      } else {
        next(`/login?redirect=${to.fullPath}`)
        NProgress.done()
      }
    }
  }
})

router.afterEach((to,from) => {
  if(to.path === '/login'){
    const menusStore = useMenusStore()
    const tabsStore = useTabsStore()
    menusStore.resetMenus()
    tabsStore.resetTabs()
  }
  NProgress.done()
})
