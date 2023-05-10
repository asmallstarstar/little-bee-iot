import { createI18n } from "vue-i18n"
import zh from "./zh"
import en from "./en"

function getNavigatorLanguage() {
    let lan = navigator.language
    if(lan.toLowerCase().indexOf('zh')!==-1){
        return 'zh'
    }else if(lan.toLowerCase().indexOf('en')!==-1){
        return 'en'
    }
    return 'en'
}

const i18n = createI18n({
    locale: localStorage.getItem("language")|| getNavigatorLanguage(),
    fallbackLocale: 'en',
    messages: {
        zh,
        en,
    },
    legacy: false,
    globalInjection: true,
})

//When you need to add a new language, just add this one
i18n.languages = [
    {
        locale:'en',
        label:"English"
    },
    {
        locale:'zh',
        label:"简体中文"
    },
]

i18n.getCurrentLanguageLocale = ()=>{
    return localStorage.getItem("language")|| getNavigatorLanguage()
}

i18n.getCurrentLanguageLabel = ()=>{
    const locale = localStorage.getItem("language")|| getNavigatorLanguage()
    const languages = i18n.languages.filter(x=>x.locale===locale)
    if(languages.length>0) {
        return languages[0].label
    }else{
        return "English"
    }
}

export default i18n
