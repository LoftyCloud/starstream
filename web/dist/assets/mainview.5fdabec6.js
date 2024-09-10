import{p as d}from"./index.8b535f41.js";import{_ as g,o as l,c as r,a as o,t as p,b as v,r as h,d as V,p as C,e as D,f as L,n as b,F as y,g as S,V as I,h as R,i as _,w as T}from"./index.eb72afd3.js";import{V as B,a as U}from"./VH-Upload.a2b28b45.js";const G={props:["message"],data(){return{visible:!1}},methods:{confirm(){this.$emit("confirm"),this.visible=!1},cancel(){this.visible=!1},show(){this.visible=!0}}},j={key:0,class:"modal-backdrop"},F={class:"modal-content"};function q(e,t,a,c,i,s){return i.visible?(l(),r("div",j,[o("div",F,[o("p",null,p(a.message),1),o("button",{class:"btn",onClick:t[0]||(t[0]=(...n)=>s.confirm&&s.confirm(...n))},"\u786E\u8BA4"),o("button",{class:"btn",onClick:t[1]||(t[1]=(...n)=>s.cancel&&s.cancel(...n))},"\u53D6\u6D88")])])):v("",!0)}var K=g(G,[["render",q],["__scopeId","data-v-28d3210f"]]);const O={name:"VH-VideoCover",components:{ConfirmationModal:K},data(){return{coverError:!0,modalMessage:"",currentAction:""}},props:{vid:null,uid:null,cover:{type:String,default:"./src/assets/img/errorcover.png"},title:{type:String,default:"\u65E0\u4E3B\u9898\u89C6\u9891"},time:{type:null,default:"00:00:00"},watch:{type:null,default:"0"},cid:{type:null,default:"\u9ED8\u8BA4"},master:{type:null,default:"admin"}},methods:{showModal(e){e==="edit"?(this.modalMessage="\u5149\u901F\u5F00\u53D1\u4E2D\uFF0C\u656C\u8BF7\u671F\u5F85",this.currentAction="edit"):e==="delete"?(this.modalMessage="\u786E\u5B9A\u8981\u5220\u9664\u8BE5\u89C6\u9891\u5417\uFF1F",this.currentAction="delete"):e==="warn"&&(this.modalMessage="\u5F53\u524D\u6743\u9650\u4E0D\u8DB3\uFF0C\u8054\u7CFB\u7AD9\u4E3B\u83B7\u53D6\u5B8C\u6574\u6743\u9650~",this.currentAction="warn"),this.$refs.confirmModal.show()},confirmAction(){this.currentAction==="delete"?this.deleteFunc():this.currentAction==="edit"?console.log("\u6267\u884C\u7F16\u8F91\u64CD\u4F5C"):this.currentAction==="warn"&&console.log("\u7528\u6237\u6743\u9650\u4E0D\u8DB3")},async deleteFunc(){let e={vid:this.vid},t={};if(await d("/admin/delete/video/information",e).then(a=>{switch(a.code){case 201:t={type:1,data:"\u5220\u9664\u9014\u4E2D\u51FA\u73B0\u5F02\u5E38"};break;case 200:t={type:2,data:"\u5220\u9664\u8BE5\u89C6\u9891\u6210\u529F"};break;default:t={type:3,data:a.msg}}}),t.type===3){this.showModal("warn");return}this.$refs.AlertMsg.addMsg(t.type,t.data),this.$emit("reset")},onClickVideo(){let e={vid:this.vid,uid:this.uid,titile:this.title,master:this.master};this.$emit("go",e)},goMaster(){let e={uid:this.uid,master:this.master};this.$emit("gomaster",e)}}},W=e=>(C("data-v-59f431c8"),e=e(),D(),e),J={class:"video-cover-box"},Q=["src"],X={key:0,class:"cover-text"},Y=W(()=>o("span",{class:"iconfont icon-tupianqueshi"},null,-1)),Z=[Y],z=["title"],tt={class:"video-info"},et={title:"\u89C6\u9891\u64AD\u653E\u91CF",class:"iconColor iconfont icon-shipin"};function it(e,t,a,c,i,s){const n=h("AlertMsg"),m=h("ConfirmationModal");return l(),r("div",J,[o("div",{onClick:t[1]||(t[1]=f=>s.onClickVideo()),class:"video-cover"},[o("img",{src:a.cover,onError:t[0]||(t[0]=f=>i.coverError=!0),alt:" "},null,40,Q),i.coverError?(l(),r("div",X,Z)):v("",!0)]),o("div",{title:a.title,class:"video-detail-box"},[o("p",{onClick:t[2]||(t[2]=f=>s.onClickVideo())},p(a.title),1)],8,z),o("div",tt,[o("span",et,p(a.watch),1),o("span",{onClick:t[3]||(t[3]=f=>s.goMaster()),title:"\u89C6\u9891\u4E0A\u4F20\u8005",class:"iconColor iconfont icon-supplier-ship"},p(a.master),1),o("span",{onClick:t[4]||(t[4]=()=>{}),title:"\u7C7B\u522B",class:"tag"},p(a.cid),1),o("span",{onClick:t[5]||(t[5]=f=>s.showModal("edit")),title:"\u7F16\u8F91\u89C6\u9891\u4FE1\u606F",class:"iconColor2 iconfont icon-wenbenshuru"}),o("span",{onClick:t[6]||(t[6]=f=>s.showModal("delete")),title:"\u5220\u9664\u89C6\u9891",class:"iconColor2 iconfont icon-shanchu"})]),V(n,{ref:"AlertMsg"},null,512),V(m,{ref:"confirmModal",message:i.modalMessage,onConfirm:s.confirmAction},null,8,["message","onConfirm"])])}var st=g(O,[["render",it],["__scopeId","data-v-59f431c8"]]);const at={name:"VH-VideoCoverSkeleton",data(){return{}},props:{},methods:{}},ot={class:"video-cover-box"},lt=L('<div class="video-cover" data-v-372aced7></div><div title="\u52A0\u8F7D\u4E2D" class="video-detail-box" data-v-372aced7><div class="skeleton-mini" data-v-372aced7></div><div class="skeleton-middle" data-v-372aced7></div></div><div class="video-info" data-v-372aced7><div class="info-skeleton" data-v-372aced7></div></div>',3),nt=[lt];function ct(e,t,a,c,i,s){return l(),r("div",ot,nt)}var dt=g(at,[["render",ct],["__scopeId","data-v-372aced7"]]);const ht={name:"VH-Empty",data(){return{}},props:{description:{type:String,default:"\u7A7A\u7684\u54E6"}},methods:{}},rt=e=>(C("data-v-be0b676a"),e=e(),D(),e),ut={class:"empty-box"},vt=rt(()=>o("span",{class:"bigBox iconfont icon-box-empty"},null,-1));function ft(e,t,a,c,i,s){return l(),r("div",ut,[vt,o("p",null,p(a.description),1)])}var _t=g(ht,[["render",ft],["__scopeId","data-v-be0b676a"]]);const mt={name:"VH-ClassSelect",data(){return{listData:[{m:"1",item:"\u53EF\u7231",key:1},{m:"2",item:"\u641E\u7B11",key:2},{m:"3",item:"\u5B66\u4E60",key:3},{m:"4",item:"\u6E38\u620F",key:4},{m:"5",item:"\u65E5\u5E38",key:5},{m:"6",item:"\u6B63\u7ECF",key:6},{m:"7",item:"\u5176\u4ED6",key:7}],select:-1}},props:{itemData:{type:Array,default:()=>[{m:"1",item:"\u53EF\u7231",key:1},{m:"2",item:"\u641E\u7B11",key:2},{m:"3",item:"\u5B66\u4E60",key:3},{m:"4",item:"\u6E38\u620F",key:4},{m:"5",item:"\u65E5\u5E38",key:5},{m:"6",item:"\u6B63\u7ECF",key:6},{m:"7",item:"\u5176\u4ED6",key:7}]},vhKey:{type:String,default:"key"},cid:{type:Number,default:0}},watch:{itemData(e){e&&e.length>0&&(this.listData=e)},cid(e){e===0&&(this.select=-1)}},created(){},methods:{onclick(e,t){if(this.select=t,e==0)this.$emit("select",0);else{let a=e.key;this.$emit("select",a)}}}},pt={class:"class-box"},gt={class:"class-inner-box"},Vt=["onClick"];function kt(e,t,a,c,i,s){return l(),r("div",pt,[o("div",gt,[o("button",{class:b(i.select==-1?"btn avatice":"btn"),onClick:t[0]||(t[0]=n=>s.onclick(0,-1))},"\u5168\u90E8",2),(l(!0),r(y,null,S(i.listData,(n,m)=>(l(),r("button",{class:b(i.select==m?"btn avatice":"btn"),key:m,onClick:f=>s.onclick(n,m)},p(n.item),11,Vt))),128))])])}var yt=g(mt,[["render",kt],["__scopeId","data-v-54bce2c1"]]);const St={name:"VH-ClassSelectSkeleton",data(){return{}},props:{},watch:{},created(){},methods:{}},Ct=e=>(C("data-v-479812a9"),e=e(),D(),e),Dt={class:"class-box"},bt=Ct(()=>o("div",{class:"class-inner-box"},[o("div",{class:"class-skeleton"})],-1)),wt=[bt];function Mt(e,t,a,c,i,s){return l(),r("div",Dt,wt)}var Ht=g(St,[["render",Mt],["__scopeId","data-v-479812a9"]]);const xt={name:"mainview",components:{VHVideoCover:st,VHVideoCoverSkeleton:dt,VHEmpty:_t,VHInfomation:B,VHUpload:U,VHPage:I,VHClassSelect:yt,VHClassSelectSkeleton:Ht},data(){return{titleName:"\u9996\u9875",isNull:!1,showDlialog:!1,dialogBox:1,videoData:[].reverse(),classData:[],allCount:1,funcEvent:0,curPage:1,apiUrl:R.baseUrl.url,isLoad:!0,funcSite:1,cid:0,isLoadClass:!1,postSite:0,search:"",itemData:[{m:"1",item:"\u53EF\u7231",key:1},{m:"2",item:"\u641E\u7B11",key:2},{m:"3",item:"\u5B66\u4E60",key:3},{m:"4",item:"\u6E38\u620F",key:4},{m:"5",item:"\u65E5\u5E38",key:5},{m:"6",item:"\u6B63\u7ECF",key:6},{m:"7",item:"\u5176\u4ED6",key:7}]}},created(){document.title="StarStream",this.getVideoData(1,21,this.cid),this.getClassData()},methods:{uploadSuc(){this.showDlialog=!1,this.getVideoData(this.curPage,21,this.cid)},goVideo(e){window.open("./video?vid="+e,"_blank")},goSearch(e){this.cid=0,e==""?(this.search="",this.curPage=1,this.funcSite==1?this.getVideoData(1,21,this.cid):this.funcSite==3&&this.getRelationVideoData(where,21,this.cid),this.funcEvent=0):(this.search=e,this.curPage=1,this.funcSite==2?this.searchVideo(1,21,this.cid,this.search):this.funcSite==4&&this.searchRelationVideo(1,21,this.cid,this.search),this.funcEvent=0)},goHeader(e){e.where=="collect"?this.goRelation():e.where=="home"?(this.cid=0,this.funcEvent=0,this.getVideoData(1,21,this.cid)):e.where=="info"?(this.dialogBox=1,this.showDlialog=!0):e.where=="upload"?(this.dialogBox=2,this.showDlialog=!0):e.where=="loginout"&&(this.logout(),console.log("\u9000\u51FA"))},logout(){localStorage.removeItem("token"),this.$store&&this.$store.dispatch("logout"),this.$router.push("/login")},goWhere(e){let t=e.where;t=="zero"?this.$refs.AlertMsg.addMsg(2,"\u641C\u7D22\u503C\u53EA\u80FD\u5927\u4E8E1"):t=="error"?this.$refs.AlertMsg.addMsg(2,"\u8F93\u5165\u503C\u6709\u8BEF"):t=="nosearch"?this.$refs.AlertMsg.addMsg(2,"\u8F93\u5165\u7684\u503C\u5927\u4E8E\u603B\u9875\u6570"):t>=1&&(this.isLoad=!0,this.funcSite==1?this.getVideoData(t,21,this.cid):this.funcSite==2?this.searchVideo(t,21,this.cid,this.search):this.funcSite==3?this.getRelationVideoData(t,21,this.cid):this.funcSite==4&&this.searchRelationVideo(t,21,this.cid,this.search),this.curPage++)},onPage(e){this.isLoad=!0,this.funcSite==1?(this.getVideoData(e,21,this.cid),this.curPage++):this.funcSite==2?(this.searchVideo(e,21,this.cid,this.search),this.curPage++):this.funcSite==3?(this.getRelationVideoData(e,21,this.cid),this.curPage++):this.funcSite==4&&(this.searchRelationVideo(e,21,this.cid,this.search),this.curPage++)},getVideoData(e,t,a){this.titleName=a===0?"\u9996\u9875":"\u5206\u7C7B",this.funcSite=1,this.videoData=[];let c={cid:a,page:e,size:t},i=a===0?"/video/list":"/video/class/list";d(i,c).then(s=>{this.videoData=s.data.list,this.funcEvent===0&&(this.allCount=Math.ceil(s.data.count/t),this.funcEvent=1),this.toTop(),!this.videoData&&this.postSite<1&&d(i,c).then(n=>{this.videoData=n.data.list,this.allCount=Math.ceil(n.data.count/t),this.toTop(),this.isLoad=!1,this.postSite=0}),this.isLoad=!1,this.isNull=!this.videoData||!this.videoData.length})},searchVideo(e,t,a,c){this.titleName="\u641C\u7D22",this.funcSite=2,this.videoData=[];let i={key:c,cid:a,page:e,size:t};d("/video/search/list",i).then(s=>{this.videoData=s.data.list,this.funcEvent==0&&(this.allCount=Math.ceil(s.data.count/t),this.funcEvent=1),this.toTop(),this.videoData==null&&this.postSite<1&&d("/video/search/list",i).then(n=>{this.videoData=n.data.list,this.allCount=Math.ceil(n.data.count/t),this.toTop(),this.isLoad=!1,this.postSite=0}),this.isLoad=!1,this.videoData[0]==null?this.isNull=!0:this.isNull=!1})},getRelationVideoData(e,t,a){this.titleName="\u6536\u85CF",this.funcSite=3,this.videoData=[];let c={cid:a,page:e,size:t};a==0?d("/relation/list",c).then(i=>{this.videoData=i.data.list,this.funcEvent==0&&(this.allCount=Math.ceil(i.data.count/t),this.funcEvent=1),this.toTop(),this.videoData==null&&this.postSite<1&&d("/relation/list",c).then(s=>{this.videoData=s.data.list,this.allCount=Math.ceil(s.data.count/t),this.toTop(),this.isLoad=!1,this.postSite=0}),this.isLoad=!1,this.videoData[0]==null?this.isNull=!0:this.isNull=!1}):d("/relation/class/list",c).then(i=>{this.videoData=i.data.list,this.funcEvent==0&&(this.allCount=Math.ceil(i.data.count/t),this.funcEvent=1),this.toTop(),this.videoData==null&&this.postSite<1&&d("/relation/class/list",c).then(s=>{this.videoData=s.data.list,this.allCount=Math.ceil(s.data.count/t),this.toTop(),this.isLoad=!1,this.postSite=0}),this.isLoad=!1,this.videoData[0]==null?this.isNull=!0:this.isNull=!1})},searchRelationVideo(e,t,a,c){this.titleName="\u6536\u85CF",this.funcSite=4,this.videoData=[];let i={key:c,cid:a,page:e,size:t};d("/relation/search/list",i).then(s=>{this.videoData=s.data.list,this.funcEvent==0&&(this.allCount=Math.ceil(s.data.count/t),this.funcEvent=1),this.toTop(),this.videoData==null&&this.postSite<1&&d("/relation/search/list",i).then(n=>{this.videoData=n.data.list,this.allCount=Math.ceil(n.data.count/t),this.toTop(),this.isLoad=!1,this.postSite=0}),this.isLoad=!1,this.videoData[0]==null?this.isNull=!0:this.isNull=!1})},goRelation(){this.cid=0,this.funcEvent=0,this.getRelationVideoData(1,21,this.cid)},getClassData(){d("/class/list","").then(e=>{this.classData=e.data,this.isLoadClass=!0,this.classData==null&&this.postSite<1&&d("/class/list","").then(t=>{this.classData=t.data,this.isLoadClass=!0,this.postSite=0})})},toTop(){document.documentElement.scrollTop=0},cilickOnClass(e){this.cid=e,this.funcSite==1?(this.curPage=1,this.funcEvent=0,this.getVideoData(1,21,this.cid)):this.funcSite==2?(this.curPage=1,this.funcEvent=0,this.searchVideo(1,21,this.cid,this.search)):this.funcSite==3?(this.curPage=1,this.funcEvent=0,this.getRelationVideoData(1,21,this.cid)):this.funcSite==4&&(this.curPage=1,this.funcEvent=0,this.searchRelationVideo(1,21,this.cid,this.search))}}},$t={class:"main-box"},Et={class:"video-list-box"},Nt={key:2,class:"bond-box"};function Pt(e,t,a,c,i,s){const n=h("Headers"),m=h("VH-ClassSelectSkeleton"),f=h("VH-ClassSelect"),w=h("VH-VideoCoverSkeleton"),M=h("VH-VideoCover"),H=h("VH-Empty"),x=h("VH-Page"),$=h("VH-Infomation"),E=h("VH-Upload"),N=h("VH-Dialog"),P=h("AlertMsg");return l(),r("div",$t,[V(n,{nameChar:i.titleName,onSearch:s.goSearch,onGo:s.goHeader},null,8,["nameChar","onSearch","onGo"]),o("div",Et,[i.isLoadClass?v("",!0):(l(),_(m,{key:0})),i.isLoadClass?(l(),_(f,{key:1,itemData:i.classData,label:"classname",vhKey:"cid",cid:i.cid,onSelect:s.cilickOnClass},null,8,["itemData","cid","onSelect"])):v("",!0),i.isLoad?(l(),r("div",Nt,[(l(),r(y,null,S(21,(u,k)=>V(w,{class:"mar-3",key:k})),64))])):v("",!0),(l(!0),r(y,null,S(i.videoData,(u,k)=>(l(),_(M,{class:"mar-3",onGo:A=>s.goVideo(u.vid),onReset:t[0]||(t[0]=A=>s.getVideoData(this.curPage,21,this.cid)),key:k,vid:u.vid,uid:u.uid,cover:i.apiUrl+"/file/cover/"+u.vid+".jpg",title:u.detail,time:u.vtime,watch:u.watch,master:u.username,cid:i.itemData[u.cid-1].item},null,8,["onGo","vid","uid","cover","title","time","watch","master","cid"]))),128)),i.isNull?(l(),_(H,{key:3,description:"\u65E0\u76F8\u5173\u89C6\u9891"})):v("",!0),i.isNull?v("",!0):(l(),_(x,{key:4,class:"page",all:i.allCount,curpage:i.curPage,onGowhere:s.goWhere,onPage:s.onPage},null,8,["all","curpage","onGowhere","onPage"]))]),i.showDlialog?(l(),_(N,{key:0,onClose:t[1]||(t[1]=u=>i.showDlialog=!1)},{default:T(()=>[i.dialogBox==1?(l(),_($,{key:0})):v("",!0),i.dialogBox==2?(l(),_(E,{key:1,onUploadSuc:s.uploadSuc},null,8,["onUploadSuc"])):v("",!0)]),_:1})):v("",!0),V(P,{ref:"AlertMsg"},null,512)])}var Rt=g(xt,[["render",Pt],["__scopeId","data-v-561816a4"]]);export{Rt as default};
//# sourceMappingURL=mainview.5fdabec6.js.map
