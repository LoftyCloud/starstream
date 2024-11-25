import{p as w}from"./index.5ab5e305.js";import{_ as F,j as _,h as x,r as v,o as d,c as a,a as s,t as m,k as b,b as u,l as g,v as h,d as M,p as k,e as y}from"./index.a591e6ed.js";const V={name:"VH-Infomation",data(){return{filename:"",useravatar:"",origin:null,selectType:1,uploadImgData:null,errForm:{oldpwd:!1,oldpwdmsg:"",newpwd:!1,newpwdmsg:"",repwd:!1,repwdmsg:""},pwdForm:{oldpwd:"",newpwd:"",repwd:""}}},props:{},created(){const o=_.get("origin");this.origin=o,this.useravatar=x.baseUrl.url+"file/avatar/"+o.avatar},methods:{uploadAvatar(){if(this.uploadImgData==null){this.$refs.AlertMsg.addMsg(2,"\u8BF7\u9009\u62E9\u4E0A\u4F20\u7684\u5934\u50CF");return}let o=new FormData;o.append("file",this.uploadImgData),w("/userinfo/upload/avatar",o).then(e=>{switch(e.code){case 201:this.$refs.AlertMsg.addMsg(2,"\u672A\u4E0A\u4F20\u6587\u4EF6");break;case 202:this.$refs.AlertMsg.addMsg(2,"\u6587\u4EF6\u5927\u4E8E2M");break;case 203:this.$refs.AlertMsg.addMsg(2,"\u6587\u4EF6\u7C7B\u578B\u4E0D\u7B26\u5408\u8981\u6C42");break;case 204:this.$refs.AlertMsg.addMsg(2,"\u5728\u5B58\u50A8\u5934\u50CF\u65F6\u51FA\u73B0\u4E86\u5F02\u5E38");break;case 200:this.$refs.AlertMsg.addMsg(1,"\u5934\u50CF\u4FEE\u6539\u6210\u529F"),this.origin.avatar=e.data.avatar,this.$cookies.set("origin",this.origin),window.open("./","_self");break;default:this.$refs.AlertMsg.addMsg(2,"\u672A\u77E5\u9519\u8BEF\uFF1A"+e.msg)}}).catch(e=>{this.$refs.AlertMsg.addMsg(2,"\u5F02\u5E38\u9519\u8BEF\uFF1A"+e)})},upload(o){let e=o.target.files[0];if(this.uploadImgData=e,e.type!=="image/png"&&e.type!=="image/jpeg"){this.$refs.AlertMsg.addMsg(2,"\u53EA\u80FD\u9009\u62E9\u56FE\u7247\u7C7B\u578B\u4E3Ajpg,png\u683C\u5F0F");return}else if(e.type==null){this.$refs.AlertMsg.addMsg(2,"\u53EA\u80FD\u9009\u62E9\u56FE\u7247\u7C7B\u578B\u4E3Ajpg,png\u683C\u5F0F");return}var p=this;p.filename=e.name;let c=new FileReader;c.readAsDataURL(e),c.addEventListener("load",function(){p.useravatar=c.result})},changePwd(){if(this.pwdForm.oldpwd.length<1)this.errForm.oldpwdmsg="\u8BF7\u8F93\u5165\u6B63\u786E\u7684\u65E7\u5BC6\u7801",this.errForm.oldpwd=!0;else if(this.pwdForm.newpwd.length<1)this.errForm.newpwdmsg="\u8BF7\u8F93\u5165\u6B63\u786E\u7684\u65B0\u5BC6\u7801",this.errForm.newpwd=!0;else if(this.pwdForm.repwd.length<1)this.errForm.repwdmsg="\u8BF7\u91CD\u590D\u65B0\u5BC6\u7801",this.errForm.repwd=!0;else if(this.pwdForm.repwd!=this.pwdForm.newpwd)this.errForm.repwdmsg="\u4E0E\u65B0\u5BC6\u7801\u4E0D\u4E00\u81F4",this.errForm.repwd=!0;else{let o={account:this.origin.account,password:this.pwdForm.oldpwd,repassword:this.pwdForm.newpwd,newpassword:this.pwdForm.repwd};w("/userinfo/update/password",o).then(e=>{switch(e.code){case 201:this.errForm.newpwdmsg="\u4E24\u6B21\u5BC6\u7801\u4E0D\u4E00\u81F4",this.errForm.newpwd=!0,this.errForm.repwdmsg="\u4E24\u6B21\u5BC6\u7801\u4E0D\u4E00\u81F4",this.errForm.repwd=!0;break;case 202:this.errForm.newpwdmsg="\u5BC6\u7801\u4E0D\u7B26\u5408\u8981\u6C42",this.errForm.newpwd=!0;break;case 203:this.$refs.AlertMsg.addMsg(2,"\u5BC6\u7801\u66F4\u6539\u65F6\u4EA7\u751F\u5F02\u5E38");break;case 200:this.$refs.AlertMsg.addMsg(1,"\u4FEE\u6539\u5BC6\u7801\u6210\u529F"),_.remove("origin"),window.open("./login","_self");break;default:this.$refs.AlertMsg.addMsg(2,"\u672A\u77E5\u9519\u8BEF\uFF1A"+e.msg)}}).catch(e=>{this.$refs.AlertMsg.addMsg(2,"\u5F02\u5E38\u9519\u8BEF\uFF1A"+e)})}}}},i=o=>(k("data-v-066d6e96"),o=o(),y(),o),D={class:"infomation-box"},C={class:"infomation-border"},$={class:"infomation-menu"},I=i(()=>s("div",{class:"user-btn-inner"},[s("span",{class:"col-right-6 iconfont icon-changyongxinxi"}),s("p",null,"\u4FEE\u6539\u5934\u50CF")],-1)),U=i(()=>s("span",{class:"iconfont icon-youxiang"},null,-1)),j=[I,U],H=i(()=>s("div",{class:"user-btn-inner"},[s("span",{class:"col-right-6 iconfont icon-mima"}),s("p",null,"\u4FEE\u6539\u5BC6\u7801")],-1)),S=i(()=>s("span",{class:"iconfont icon-youxiang"},null,-1)),T=[H,S],q={class:"infomation-context"},N={key:0,class:"avatar-box"},P=["src"],B={class:"file-box"},R={href:"javascript:;",class:"file"},E={key:1,class:"password-box"},L={class:"pwd-b"},z={class:"user-box"},G=i(()=>s("label",null,"\u65E7\u5BC6\u7801",-1)),J={key:0},K=i(()=>s("i",{class:"col-ri-5 iconfont icon-cuowu1"},null,-1)),O={class:"user-box"},Q=i(()=>s("label",null,"\u65B0\u5BC6\u7801",-1)),W={key:0},X=i(()=>s("i",{class:"col-ri-5 iconfont icon-cuowu1"},null,-1)),Y={class:"user-box"},Z=i(()=>s("label",null,"\u91CD\u590D\u65B0\u5BC6\u7801",-1)),ee={key:0},se=i(()=>s("i",{class:"col-ri-5 iconfont icon-cuowu1"},null,-1));function te(o,e,p,c,t,l){const f=v("AlertMsg");return d(),a("div",D,[s("div",C,[s("div",$,[s("a",{onClick:e[0]||(e[0]=r=>t.selectType=1),class:"user-btn"},j),s("a",{onClick:e[1]||(e[1]=r=>t.selectType=2),class:"user-btn"},T)]),s("div",q,[t.selectType==1?(d(),a("div",N,[s("img",{src:t.useravatar},null,8,P),s("div",B,[s("p",null,m(t.filename),1),s("a",R,[b("\u9009\u62E9\u6587\u4EF6 "),s("input",{type:"file",id:"file",onChange:e[2]||(e[2]=(...r)=>l.upload&&l.upload(...r)),accept:".jpg,.png,.jpeg"},null,32)])]),s("button",{class:"btn",onClick:e[3]||(e[3]=(...r)=>l.uploadAvatar&&l.uploadAvatar(...r))},"\u786E\u8BA4\u4FEE\u6539")])):u("",!0),t.selectType==2?(d(),a("div",E,[s("div",L,[s("form",null,[s("div",z,[g(s("input",{type:"password","onUpdate:modelValue":e[4]||(e[4]=r=>t.pwdForm.oldpwd=r),required:"",autocomplete:"on",onChange:e[5]||(e[5]=r=>t.errForm.oldpwd=!1)},null,544),[[h,t.pwdForm.oldpwd]]),G,t.errForm.oldpwd?(d(),a("span",J,[K,s("p",null,m(t.errForm.oldpwdmsg),1)])):u("",!0)]),s("div",O,[g(s("input",{type:"password","onUpdate:modelValue":e[6]||(e[6]=r=>t.pwdForm.newpwd=r),required:"",autocomplete:"on",onChange:e[7]||(e[7]=r=>t.errForm.newpwd=!1)},null,544),[[h,t.pwdForm.newpwd]]),Q,t.errForm.newpwd?(d(),a("span",W,[X,s("p",null,m(t.errForm.newpwdmsg),1)])):u("",!0)]),s("div",Y,[g(s("input",{type:"password","onUpdate:modelValue":e[8]||(e[8]=r=>t.pwdForm.repwd=r),required:"",autocomplete:"on",onChange:e[9]||(e[9]=r=>t.errForm.repwd=!1)},null,544),[[h,t.pwdForm.repwd]]),Z,t.errForm.repwd?(d(),a("span",ee,[se,s("p",null,m(t.errForm.repwdmsg),1)])):u("",!0)])]),s("button",{class:"btn-m",onClick:e[10]||(e[10]=(...r)=>l.changePwd&&l.changePwd(...r))},"\u786E\u8BA4\u4FEE\u6539")])])):u("",!0)])]),M(f,{ref:"AlertMsg"},null,512)])}var we=F(V,[["render",te],["__scopeId","data-v-066d6e96"]]);const oe={name:"VH-Upload",data(){return{uploadForm:{title:"",class:null},errForm:{title:!1,titlemsg:""},select:"",curData:"\u9009\u62E9\u89C6\u9891\u5206\u533A",itemData:[{m:"1",item:"\u53EF\u7231",key:1},{m:"2",item:"\u641E\u7B11",key:2},{m:"3",item:"\u5B66\u4E60",key:3},{m:"4",item:"\u6E38\u620F",key:4},{m:"5",item:"\u65E5\u5E38",key:5},{m:"6",item:"\u6B63\u7ECF",key:6},{m:"7",item:"\u5176\u4ED6",key:7}],filename:""}},props:{},created(){const o=_.get("origin");this.origin=o},methods:{selectClass(o){this.uploadForm.class=o},upload(o){console.log(o.target.files[0]);let e=o.target.files[0];if(e.type!=="video/mp4"){this.$refs.AlertMsg.addMsg(2,"\u53EA\u80FD\u9009\u62E9\u89C6\u9891\u7C7B\u578B\u4E3Amp4\u683C\u5F0F");return}else if(e.type==null){this.$refs.AlertMsg.addMsg(2,"\u53EA\u80FD\u9009\u62E9\u89C6\u9891\u7C7B\u578B\u4E3Amp4\u683C\u5F0F");return}var p=this;p.filename=e.name,this.filename=e.name,this.uploadForm.file=e,this.uploadForm.title=e.name.replace(/\.[^/.]+$/,"")},async uploadVideo(){if(this.uploadForm.title===""){this.errForm.title=!0,this.errForm.titlemsg="\u89C6\u9891\u6807\u9898\u4E0D\u80FD\u4E3A\u7A7A",console.log("\u89C6\u9891\u6807\u9898\u4E0D\u80FD\u4E3A\u7A7A");return}if(!this.uploadForm.file){this.$refs.AlertMsg.addMsg(2,"\u8BF7\u5148\u9009\u62E9\u4E00\u4E2A\u89C6\u9891\u6587\u4EF6"),console.log("\u8BF7\u5148\u9009\u62E9\u4E00\u4E2A\u89C6\u9891\u6587\u4EF6");return}if(!this.uploadForm.class){this.$refs.AlertMsg.addMsg(2,"\u8BF7\u9009\u62E9\u4E00\u4E2A\u5206\u7C7B"),console.log("\u8BF7\u9009\u62E9\u4E00\u4E2A\u5206\u7C7B");return}let o=new FormData;o.append("detail",this.uploadForm.title),o.append("file",this.uploadForm.file),o.append("cid",this.uploadForm.class),w("/video/upload",o).then(e=>{switch(console.log("\u51C6\u5907\u4E0A\u4F20\u89C6\u9891..."),e.code){case 602:this.$refs.AlertMsg.addMsg(2,"cid\u57FA\u672C\u53C2\u6570\u9519\u8BEF");break;case 600:this.$refs.AlertMsg.addMsg(2,"\u6587\u4EF6\u53C2\u6570\u9519\u8BEF");break;case 601:this.$refs.AlertMsg.addMsg(2,"detail\u57FA\u672C\u53C2\u6570\u9519\u8BEF");break;case 201:this.$refs.AlertMsg.addMsg(2,"\u6587\u4EF6\u7C7B\u578B\u4E0D\u7B26\u5408\u8981\u6C42");break;case 202:this.$refs.AlertMsg.addMsg(2,"\u8BFB\u53D6\u6587\u4EF6\u6D41\u5931\u8D25");break;case 203:this.$refs.AlertMsg.addMsg(2,"\u89C6\u9891\u4E0A\u4F20\u51FA\u73B0\u95EE\u9898");break;case 204:this.$refs.AlertMsg.addMsg(2,"\u89C6\u9891\u5B58\u50A8\u8FC7\u7A0B\u51FA\u73B0\u95EE\u9898");break;case 200:this.$refs.AlertMsg.addMsg(1,"\u89C6\u9891\u4E0A\u4F20\u6210\u529F"),this.uploadForm.title="",this.uploadForm.class=null,this.filename="",this.$emit("uploadSuc");break;default:this.$refs.AlertMsg.addMsg(2,"\u672A\u77E5\u9519\u8BEF\uFF1A"+e.msg)}}).catch(e=>{this.$refs.AlertMsg.addMsg(2,"\u5F02\u5E38\u9519\u8BEF\uFF1A"+e)})}}},A=o=>(k("data-v-6d95b5dd"),o=o(),y(),o),re={class:"upload-box"},le={class:"upload-inner"},ie={class:"top-bard"},ne={class:"user-box"},de=A(()=>s("label",null,"\u89C6\u9891\u6807\u9898",-1)),ae={key:0},pe=A(()=>s("i",{class:"col-ri-5 iconfont icon-cuowu1"},null,-1)),ue={class:"video-area"},me={class:"file-box"},ce={href:"javascript:;",class:"file"};function ge(o,e,p,c,t,l){const f=v("VH-Select"),r=v("AlertMsg");return d(),a("div",re,[s("div",le,[s("form",ie,[s("div",ne,[g(s("input",{type:"text","onUpdate:modelValue":e[0]||(e[0]=n=>t.uploadForm.title=n),required:"",autocomplete:"on",onChange:e[1]||(e[1]=n=>t.errForm.title=!1)},null,544),[[h,t.uploadForm.title]]),de,t.errForm.title?(d(),a("span",ae,[pe,s("p",null,m(t.errForm.titlemsg),1)])):u("",!0)]),s("div",ue,[M(f,{width:"170",onChange:l.selectClass,modelValue:t.select,"onUpdate:modelValue":e[2]||(e[2]=n=>t.select=n),curData:t.curData,itemData:t.itemData,label:"item","vh-key":"m"},null,8,["onChange","modelValue","curData","itemData"])])]),s("div",null,[s("div",me,[s("p",null,m(t.filename),1),s("a",ce,[b("\u9009\u62E9\u6587\u4EF6 "),s("input",{type:"file",id:"file",onChange:e[3]||(e[3]=(...n)=>l.upload&&l.upload(...n)),accept:".mp4"},null,32)])])]),s("button",{class:"btn-m",onClick:e[4]||(e[4]=(...n)=>l.uploadVideo&&l.uploadVideo(...n))},"\u4E0A\u4F20\u89C6\u9891")]),M(r,{ref:"AlertMsg"},null,512)])}var _e=F(oe,[["render",ge],["__scopeId","data-v-6d95b5dd"]]);export{we as V,_e as a};
//# sourceMappingURL=VH-Upload.ec015234.js.map