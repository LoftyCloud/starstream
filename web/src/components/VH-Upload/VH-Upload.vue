<template>
    <div class="upload-box">
        <div class="upload-inner">
            <form class="top-bard">
                <div class="user-box">
                    <input type="text"  v-model="uploadForm.title" required autocomplete="on" @change="errForm.title=false">
                    <label>视频标题</label>
                    <span v-if="errForm.title"><i class="col-ri-5 iconfont icon-cuowu1"></i><p>{{errForm.titlemsg}}</p></span>
                </div>
                <div class="video-area">
                    <VH-Select width="170" @change="selectClass" v-model="select" :curData="curData" :itemData="itemData" label="item" vh-key="m"></VH-Select>
                </div>
            </form>
            <div>
                <div class="file-box">
                    <p>{{filename}}</p>
                    <a href="javascript:;" class="file">选择文件
                        <input type="file" id="file" @change="upload" accept=".mp4">
                    </a>
                </div>
            </div>
            <button class="btn-m" @click="uploadVideo">上传视频</button>
        </div>
        <AlertMsg ref="AlertMsg"></AlertMsg>
    </div>
</template>

<script>
import cookies from "vue-cookies";
import { postRequest } from '@/api';
export default {
    name: 'VH-Upload',
    data () {
        return {
            uploadForm:{
                title:'',
                class:null,
            },
            errForm:{
                title:false,
                titlemsg:'',
            },
            select:'',
			curData: '选择视频分区',
			itemData: [
          {m:'1',item:'可爱',key:1},
          {m:'2',item:'搞笑',key:2},
          {m:'3',item:'学习',key:3},
          {m:'4',item:'游戏',key:4},
          {m:'5',item:'日常',key:5},
          {m:'6',item:'正经',key:6},
          {m:'7',item:'其他',key:7}],
          filename:'',
        }
    },
    props:{
        
    },
    created() {
        const origin = cookies.get('origin');
        this.origin=origin;
    },
    methods:{
        selectClass(e){
            this.uploadForm.class = e
            // console.log(this.uploadForm.class)
        },

        upload(e){
            //e.target指向事件执行时鼠标所点击区域的那个元素，那么为input的标签，
            // 可以输出 e.target.files 看看，这个files对象保存着选中的所有的图片的信息。
            console.log(e.target.files[0])
            //------------------------------------------------------     
            // 既然如此，循环这个files对象，用for of 循环，     
            let item = e.target.files[0]
            //正则表达式，判断每个元素的type属性是否为图片形式，如图
            if (item.type !== 'video/mp4') {
                // 提示只能是图片，return
                this.$refs.AlertMsg.addMsg(
                    2,"只能选择视频类型为mp4格式"
                )
                return;
            }else if(item.type==undefined){
                this.$refs.AlertMsg.addMsg(
					          2,"只能选择视频类型为mp4格式"
				        )
                return;
            }

            // 保存下当前 this ，就是vue实例
            var _this = this;
            _this.filename = item.name

            this.filename = item.name
            this.uploadForm.file = item  // 将文件保存到uploadForm中，方便后续使用
            this.uploadForm.title = item.name.replace(/\.[^/.]+$/, '');  // 去掉最后一个点和之后的部分
            ;
        },


        async uploadVideo(){
            // 1. 验证表单数据
            if (this.uploadForm.title === '') {
                this.errForm.title = true
                this.errForm.titlemsg = '视频标题不能为空'
                console.log('视频标题不能为空')
                return;
            }

            if (!this.uploadForm.file) {
                this.$refs.AlertMsg.addMsg(2, "请先选择一个视频文件")
                console.log('请先选择一个视频文件');
                return
            }

            if (!this.uploadForm.class) {
                this.$refs.AlertMsg.addMsg(2, "请选择一个分类")
                console.log('请选择一个分类');
                return
            }

            // 2. 创建FormData对象并添加数据
            let formData = new FormData()
            formData.append('detail', this.uploadForm.title)
            formData.append('file', this.uploadForm.file)  // 添加视频文件
            formData.append('cid', this.uploadForm.class);

            // 3. 发送上传请求（使用axios）
            postRequest('/video/upload',formData).then(res => {
                console.log('准备上传视频...');
                switch (res.code) {
                    case 602:
                        this.$refs.AlertMsg.addMsg(
                            2,"cid基本参数错误"
                        )
                        break
                    case 600:
                        this.$refs.AlertMsg.addMsg(
                            2,"文件参数错误"
                        )
                        break
                    case 601:
                        this.$refs.AlertMsg.addMsg(
                            2,"detail基本参数错误"
                        )
                        break
                    case 201:
                        this.$refs.AlertMsg.addMsg(
                            2,"文件类型不符合要求"
                        )
                        break
                    case 202:
                        this.$refs.AlertMsg.addMsg(
                            2,"读取文件流失败"
                        )
                        break
                    case 203:
                        this.$refs.AlertMsg.addMsg(
                            2,"视频上传出现问题"
                        )
                        break
                    case 204:
                        this.$refs.AlertMsg.addMsg(
                            2,"视频存储过程出现问题"
                        )
                        break
                    case 200:
                        this.$refs.AlertMsg.addMsg(
                            1,"视频上传成功"
                        )
                        // 清空表单
                        this.uploadForm.title = ''
                        this.uploadForm.class = null
                        this.filename = ''
                        // 关闭窗口并刷新主页面
                        this.$emit("uploadSuc")
                        break
                    default:
                        this.$refs.AlertMsg.addMsg(
                            2,"未知错误："+res.msg
                        )
                }
            }).catch(err => {
                this.$refs.AlertMsg.addMsg(
                    2,"异常错误："+err
                )
            })
        },
    },
}

</script>

<style  scoped>
*{
    margin: 0;
    padding: 0;
}
video{
    width: 480px;
    height: 300px;
}
.upload-box{
    width: 100%;
    height: 100%;
    padding: 30px;
    box-sizing: border-box;
    background-color: #2D3047;
    border-radius: 8px;
}


.upload-inner{
    width: 100%;
    height: 100%;
    background-color: #2D3047;
    border-radius: 6px;
    //box-shadow: 0 3px 12px rgb(0 36 153 / 6%);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    box-sizing: border-box;
    animation: showBoxUp .5s;
    user-select: none;
}
.user-box{
    position: relative;
    position: relative;
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 250px;
}
.user-box input {
    width: 100%;
    padding: 10px 0;
    font-size: 16px;
    color: #FFFFFF;
    margin-bottom: 30px;
    border: none;
    border-bottom: 3px solid #1C1C1E;
    outline: none;
    background: transparent;
}
.user-box label {
    position: absolute;
    top:0;
    left: 0;
    padding: 10px 0;
    font-size: 16px;
    color: #F3CA40;
    pointer-events: none;
    transition: .5s;
}
.user-box span {
    position: absolute;
    top:0;
    right: 0;
    padding: 10px 0;
    font-size: 10px;
    color: #F3CA40;
    pointer-events: none;
    transition: .5s;
    display: flex;
    align-items: center;
}

.user-box input:focus ~ label,
.user-box input:valid ~ label {
    top: -20px;
    left: 0;
    color: #F3CA40;
    font-size: 12px;
}
.top-bard{
    display: flex;
    flex-direction: row;
    align-items: center;
}
@keyframes showBoxUp{
	0%{
		opacity: 0;
		transform:translateY(-50px);
	}
	100%{
		opacity: 1;
		transform:translateY(0px);
	}
}
.video-area{
    width: auto;
    padding: 0px 0px 30px;
    box-sizing: border-box;
    margin-left: 25px;
}
.file-box{
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-top: 30px;
}
.file-box p{
    margin-right: 10px;
}
.file {
    position: relative;
    display: inline-block;
    background-color: #9C89B8;
    border: 3px solid #1C1C1E;
    color: #fff;
    border-radius: 4px;
    padding: 4px 12px;
    overflow: hidden;
    text-decoration: none;
    text-indent: 0;
    line-height: 30px;
}
.file input {
    position: absolute;
    font-size: 100px;
    right: 0;
    top: 0;
    opacity: 0;
}
.file:hover {
    background: #F3CA40;
    border-color: #2D3047;
    color: #000000;
    text-decoration: none;
}
#file{
    cursor: pointer;
}
.btn-m{
    margin-top: 30px;
    margin-left: 11px;
    width: 90px;
    height: 40px;
    background-color: #9C89B8;
    border: 3px solid #1C1C1E;
    color: #fff;
    border-radius: 4px;
    cursor: pointer;
    font-size: 15px;
}
.btn-m:hover{
    background: #F3CA40;
    border-color: #2D3047;
    color: #000000;
}
</style>
