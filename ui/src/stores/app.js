import { useStorage } from "@vueuse/core"

const initalState = {
    token: undefined
}

export default useStorage('vblog', initalState, localStorage, {
    mergeDefaults: true
})