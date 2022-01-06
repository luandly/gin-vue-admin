import { login, getUserInfo, setUserInfo as setUserInfoApi } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouterStore } from './router'

export const useUserStore = defineStore('user', () => {
  const routerStore = useRouterStore()

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {},
    sideMode: 'dark',
    activeColor: '#4D70FF',
    baseColor: '#fff'
  })
  const token = ref('')

  const setUserInfo = (val) => {
    userInfo.value = val
  }

  const setToken = (val) => {
    token.value = val
  }

  const NeedInit = () => {
    token.value = ''
    sessionStorage.clear()
    router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }

  const GetUserInfo = async() => {
    const res = await getUserInfo()
    if (res.code === 0) {
      setUserInfo(res.data.userInfo)
    }
    return res
  }

  const LoginIn = async(loginInfo) => {
    const res = await login(loginInfo)
    if (res.code === 0) {
      setUserInfo(res.data.user)
      setToken(res.data.token)
      await routerStore.SetAsyncRouter({})
      const asyncRouters = routerStore.asyncRouters
      asyncRouters.forEach(asyncRouter => {
        router.addRoute(asyncRouter)
      })
      router.push({ name: userInfo.value.authority.defaultRouter })
      return true
    }
  }

  const LoginOut = async() => {
    const res = await jsonInBlacklist()
    if (res.code === 0) {
      token.value = ''
      sessionStorage.clear()
      router.push({ name: 'Login', replace: true })
      window.location.reload()
    }
  }

  const changeSideMode = async(data) => {
    const res = await setUserInfoApi({ sideMode: data, ID: userInfo.value.ID })
    if (res.code === 0) {
      userInfo.value.sideMode = data
      ElMessage({
        type: 'success',
        message: '设置成功'
      })
    }
  }

  const mode = computed(() => userInfo.value.sideMode)
  const sideMode = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#191a23'
    } else if (userInfo.value.sideMode === 'light') {
      return '#fff'
    } else {
      return userInfo.value.sideMode
    }
  })
  const baseColor = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#fff'
    } else if (userInfo.value.sideMode === 'light') {
      return '#191a23'
    } else {
      return userInfo.value.baseColor
    }
  })
  const activeColor = computed(() => {
    if (userInfo.value.sideMode === 'dark' || userInfo.value.sideMode === 'light') {
      return '#4D70FF'
    }
    return userInfo.activeColor
  })

  return {
    userInfo,
    token,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    changeSideMode,
    mode,
    sideMode,
    setToken,
    baseColor,
    activeColor
  }
})
