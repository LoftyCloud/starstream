import { createRouter, createWebHistory } from "vue-router";
const mainview = () => import ('../views/mainview/mainview.vue');
const login = () => import ('../views/login/login.vue');
const video = () => import ('../views/video/video.vue');

const routes = [
	{ path: '/mainview', name: 'mainview', component: mainview },
	{ path: '/login', name: 'login', component: login },
	{ path: '/register', name: 'register', component: login },
	{ path: '/video', name: 'video', component: video },
	{ path: '/:pathMatch(.*)' , name: '404', redirect: '/mainview' },
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
