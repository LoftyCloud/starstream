<template>
	<div class="main-box">
		<Headers :nameChar="titleName" @search="goSearch" @go="goHeader"></Headers>
		<div class="video-list-box">

			<VH-ClassSelectSkeleton v-if="!isLoadClass"></VH-ClassSelectSkeleton>
			<VH-ClassSelect  v-if="isLoadClass" :itemData="classData" label="classname" vhKey="cid" :cid="cid" @select="cilickOnClass"></VH-ClassSelect>

			<div class="bond-box" v-if="isLoad">
				<VH-VideoCoverSkeleton class="mar-3" v-for="(_,k) in 21" :key="k"></VH-VideoCoverSkeleton>
			</div>

<!--			基于videoData循环添加video-->
			<VH-VideoCover class="mar-3" @go="goVideo(item.vid)" @reset="getVideoData(this.curPage,21,this.cid)" v-for="(item,key) in videoData" :key="key" :vid="item.vid"
										 :uid="item.uid" :cover="apiUrl+'file/cover/'+item.vid+'.jpg'"
										 :title="item.detail" :time="item.vtime" :watch="item.watch" :master="item.username" :cid="itemData[item.cid-1].item">
			</VH-VideoCover>

			<VH-Empty v-if="isNull" description="无相关视频"></VH-Empty>
			<VH-Page v-if="!isNull" class="page" :all="allCount" :curpage="curPage" @gowhere="goWhere" @page="onPage"></VH-Page>

		</div>
		<VH-Dialog v-if="showDlialog" @close="showDlialog=false">
			<VH-Infomation v-if="dialogBox==1"></VH-Infomation>
			<VH-Upload @uploadSuc="uploadSuc" v-if="dialogBox==2"></VH-Upload>
		</VH-Dialog>

		<AlertMsg ref="AlertMsg"></AlertMsg>
	</div>
</template>

<script>
import { postRequest } from '@/api';
import VHVideoCover from '@/components/VH-VideoCover/VH-VideoCover.vue';
import VHVideoCoverSkeleton from '@/components/VH-VideoCoverSkeleton/VH-VideoCoverSkeleton.vue';
import VHEmpty from '@/components/VH-Empty/VH-Empty.vue';
import VHInfomation from '@/components/VH-Infomation/VH-Infomation.vue';
import VHUpload from '@/components/VH-Upload/VH-Upload.vue';
import VHPage from '@/components/VH-Page/VH-Page.vue';
import VHClassSelect from '@/components/VH-ClassSelect/VH-ClassSelect.vue';
import VHClassSelectSkeleton from '@/components/VH-ClassSelectSkeleton/VH-ClassSelectSkeleton.vue';
import config from '@/api/config';
export default {
	name: 'mainview',
	components:{
		VHVideoCover,
		VHVideoCoverSkeleton,
		VHEmpty,
		VHInfomation,
		VHUpload,
		VHPage,
		VHClassSelect,
		VHClassSelectSkeleton,
	},
	data() {
		return {
			titleName:'首页',
			isNull:false,
			showDlialog:false,
			dialogBox:1,  // 用于控制显示修改用户信息和上传视频框
			videoData:[].reverse(),
			classData:[],
			allCount:1,
			funcEvent:0,
			curPage:1,
			apiUrl : config.baseUrl.url,
			isLoad:true,
			funcSite:1,
			cid:0,
			isLoadClass:false,
			postSite:0,
			search:'',
			itemData: [
				{m:'1',item:'可爱',key:1},
				{m:'2',item:'搞笑',key:2},
				{m:'3',item:'学习',key:3},
				{m:'4',item:'游戏',key:4},
				{m:'5',item:'日常',key:5},
				{m:'6',item:'正经',key:6},
				{m:'7',item:'其他',key:7}],
		};
	},
	created() {
		document.title ='StarStream';
		this.getVideoData(1,21,this.cid);
		this.getClassData();
	},
	methods: {
		uploadSuc(){
			this.showDlialog = false
			this.getVideoData(this.curPage,21,this.cid)
		},

		goVideo(json){
			console.log(location.origin);
			window.open(location.origin + '/video?vid='+json, '_blank')
		},

		goSearch(search){
			this.cid = 0
			if(search==''){
				this.search=''
				this.curPage=1
				if(this.funcSite == 1){
					this.getVideoData(1,21,this.cid)
				}else if(this.funcSite == 3){
					this.getRelationVideoData(where,21,this.cid)
				}
				this.funcEvent = 0
			}else{
				this.search = search
				this.curPage=1
				if(this.funcSite == 2){
					this.searchVideo(1,21,this.cid,this.search)
				}else if(this.funcSite == 4){
					this.searchRelationVideo(1,21,this.cid,this.search)
				}
				this.funcEvent = 0
			}
		},
		goHeader(e){
			if(e.where=='collect'){
				this.goRelation();
			}else if(e.where=='home'){
				this.cid=0
				this.funcEvent = 0
				this.getVideoData(1,21,this.cid);
			}else if(e.where == 'info'){
				this.dialogBox = 1
				this.showDlialog=true
			}else if(e.where == 'upload'){
				this.dialogBox = 2
				this.showDlialog=true
			}else if(e.where == 'loginout'){
				this.logout();
				console.log("退出")
			}
		},
		logout() {
			// 1. 清除本地存储的令牌（例如 JWT）
			localStorage.removeItem('token'); // 根据你的实际存储键值进行调整
			// 2. 如果使用 Vuex，清空 Vuex 状态
			if (this.$store) {
				this.$store.dispatch('logout'); // 确保在 Vuex 中有对应的 logout 动作
			}
			// 3. 清除 Vuex 中的用户信息（可选）
			// this.$store.commit('CLEAR_USER'); // 根据你的 Vuex 配置调整
			// 4. 重定向到登录页面
			this.$router.push('/login'); // 确保 '/login' 是你的登录页面路由
			// 5. 输出日志（可选）
			// console.log("已成功退出登录");
		},
		goWhere(e){
			let where = e.where
			if(where == 'zero'){
				this.$refs.AlertMsg.addMsg(
						2,"搜索值只能大于1"
				)
			}else if(where == 'error'){
				this.$refs.AlertMsg.addMsg(
						2,"输入值有误"
				)
			}else if(where =='nosearch'){
				this.$refs.AlertMsg.addMsg(
						2,"输入的值大于总页数"
				)
			}else if(where >= 1 ){
				this.isLoad=true
				if(this.funcSite == 1){
					this.getVideoData(where,21,this.cid)
				}else if(this.funcSite == 2){
					this.searchVideo(where,21,this.cid,this.search)
				}else if(this.funcSite == 3){
					this.getRelationVideoData(where,21,this.cid)
				}else if(this.funcSite == 4){
					this.searchRelationVideo(where,21,this.cid,this.search)
				}
				this.curPage++
			}
		},
		onPage(e){
			this.isLoad=true
			if(this.funcSite == 1){
				this.getVideoData(e,21,this.cid)
				this.curPage++
			}else if(this.funcSite == 2){
				this.searchVideo(e,21,this.cid,this.search)
				this.curPage++
			}else if(this.funcSite == 3){
				this.getRelationVideoData(e,21,this.cid)
				this.curPage++
			}else if(this.funcSite == 4){
				this.searchRelationVideo(e,21,this.cid,this.search)
				this.curPage++
			}
		},

		getVideoData(page, size, cid) {
			this.titleName = cid === 0 ? '首页' : '分类'; // 根据 cid 区分标题
			this.funcSite = 1;
			this.videoData = [];
			let params = {
				"cid": cid,
				"page": page,
				"size": size
			};

			// 确定请求的 URL
			let url = cid === 0 ? '/video/list' : '/video/class/list';

			// 执行请求
			postRequest(url, params).then(res => {
				this.videoData = res.data.list;

				// 只在第一次加载时计算页数
				if (this.funcEvent === 0) {
					this.allCount = Math.ceil(res.data.count / size);
					this.funcEvent = 1;
				}

				this.toTop(); // 滚动到页面顶部

				// 如果没有数据且允许多次加载，则再次尝试加载数据
				if (!this.videoData && this.postSite < 1) {
					postRequest(url, params).then(res => {
						this.videoData = res.data.list;
						this.allCount = Math.ceil(res.data.count / size);
						this.toTop();
						this.isLoad = false;
						this.postSite = 0;
					});
				}

				this.isLoad = false;
				this.isNull = !this.videoData || !this.videoData.length;
			});
		},
		searchVideo(page,size,cid,search){
			this.titleName='搜索'
			this.funcSite = 2
			this.videoData = []
			let params = {
				"key": search,
				"cid": cid,
				"page": page,
				"size": size
			}
			postRequest('/video/search/list',params).then(res => {
				this.videoData = res.data.list
				if(this.funcEvent==0){
					this.allCount = Math.ceil(res.data.count / size)
					this.funcEvent = 1
				}
				this.toTop()
				if(this.videoData==null && this.postSite < 1){
					postRequest('/video/search/list',params).then(res => {
						this.videoData = res.data.list
						this.allCount = Math.ceil(res.data.count / size)
						this.toTop()
						this.isLoad=false
						this.postSite=0
					})
				}
				this.isLoad=false
				if(this.videoData[0]==null){
					this.isNull=true
				}else{
					this.isNull=false
				}
			})
		},

		getRelationVideoData(page,size,cid){
			this.titleName='收藏'
			this.funcSite=3
			this.videoData = []
			let params = {
				"cid": cid,
				"page": page,
				"size": size
			}
			if(cid == 0){
				postRequest('/relation/list',params).then(res => {
					this.videoData = res.data.list
					if(this.funcEvent==0){
						this.allCount = Math.ceil(res.data.count / size)
						this.funcEvent = 1
					}
					this.toTop()
					if(this.videoData==null && this.postSite < 1){
						postRequest('/relation/list',params).then(res => {
							this.videoData = res.data.list
							this.allCount = Math.ceil(res.data.count / size)
							this.toTop()
							this.isLoad=false
							this.postSite=0
						})
					}
					this.isLoad=false
					if(this.videoData[0]==null){
						this.isNull=true
					}else{
						this.isNull=false
					}
				})
			}else{
				postRequest('/relation/class/list',params).then(res => {
					this.videoData = res.data.list
					if(this.funcEvent==0){
						this.allCount = Math.ceil(res.data.count / size)
						this.funcEvent = 1
					}
					this.toTop()
					if(this.videoData==null && this.postSite < 1){
						postRequest('/relation/class/list',params).then(res => {
							this.videoData = res.data.list
							this.allCount = Math.ceil(res.data.count / size)
							this.toTop()
							this.isLoad=false
							this.postSite=0
						})
					}
					this.isLoad=false
					if(this.videoData[0]==null){
						this.isNull=true
					}else{
						this.isNull=false
					}
				})
			}
		},
		searchRelationVideo(page,size,cid,search){
			this.titleName='收藏'
			this.funcSite = 4
			this.videoData = []
			let params = {
				"key": search,
				"cid": cid,
				"page": page,
				"size": size
			}
			postRequest('/relation/search/list',params).then(res => {
				this.videoData = res.data.list
				if(this.funcEvent==0){
					this.allCount = Math.ceil(res.data.count / size)
					this.funcEvent = 1
				}
				this.toTop()
				if(this.videoData==null && this.postSite < 1){
					postRequest('/relation/search/list',params).then(res => {
						this.videoData = res.data.list
						this.allCount = Math.ceil(res.data.count / size)
						this.toTop()
						this.isLoad=false
						this.postSite=0
					})
				}
				this.isLoad=false
				if(this.videoData[0]==null){
					this.isNull=true
				}else{
					this.isNull=false
				}
			})
		},

		goRelation(){
			this.cid=0
			this.funcEvent = 0
			this.getRelationVideoData(1,21,this.cid)
		},

		getClassData(){
			postRequest('/class/list','').then(res => {
				this.classData = res.data;
				this.isLoadClass=true
				if(this.classData==null && this.postSite < 1){
					postRequest('/class/list','').then(res => {
						this.classData = res.data;
						this.isLoadClass=true
						this.postSite=0
					})
				}
			})
		},

		toTop() {
			document.documentElement.scrollTop = 0;
		},

		cilickOnClass(e){
			this.cid=e
			// console.log(this.funcSite)
			if(this.funcSite == 1){
				this.curPage=1
				this.funcEvent = 0
				this.getVideoData(1,21,this.cid)
			}else if(this.funcSite == 2){
				this.curPage=1
				this.funcEvent = 0
				this.searchVideo(1,21,this.cid,this.search)
			}else if(this.funcSite == 3){
				this.curPage=1
				this.funcEvent = 0
				this.getRelationVideoData(1,21,this.cid)
			}else if(this.funcSite == 4){
				this.curPage=1
				this.funcEvent = 0
				this.searchRelationVideo(1,21,this.cid,this.search)
			}
		},
	}
}
</script>

<style scoped>
@import '@/assets/css/mainview.css';
</style>