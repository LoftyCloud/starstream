<template>
    <div class="video-cover-box">

        <div  @click="onClickVideo()" class="video-cover">
            <img :src="cover" @error="coverError = true" alt=" ">
            <div v-if="coverError" class="cover-text">
                <span class="iconfont icon-tupianqueshi"></span>
            </div>
        </div>

        <div :title="title" class="video-detail-box">
            <p  @click="onClickVideo()">{{title}}</p>
        </div>

        <div class="video-info">
            <span title="视频播放量" class="iconColor iconfont icon-shipin">{{watch}}</span>
            <span @click="goMaster()" title="视频上传者" class="iconColor iconfont icon-supplier-ship">{{master}}</span>

            <span @click="" title="类别" class="tag">{{cid}}</span>

            <span @click="showModal('edit')" title="编辑视频信息" class="iconColor2 iconfont icon-wenbenshuru"></span>
            <span @click="showModal('delete')" title="删除视频" class="iconColor2 iconfont icon-shanchu"></span>


        </div>
        <AlertMsg ref="AlertMsg"></AlertMsg>
        <ConfirmationModal ref="confirmModal" :message="modalMessage" @confirm="confirmAction"></ConfirmationModal>
    </div>
</template>

<script>
import ConfirmationModal from '@/components/ConfirmationModal/ConfirmationModal.vue';  // 引入模态框组件
import { postRequest } from '@/api';
export default {
    name: "VH-VideoCover",

    components: {
        ConfirmationModal
    },
    data(){
        return {
            // cover: '/path/to/cover.jpg',
            coverError: true, // 图像加载错误状态
            modalMessage: '',
            currentAction: '',  // 用来区分操作的变量
        };
    },
    props:{
        vid:null,
        uid:null,
        cover:{
			type: String,
			default: './src/assets/img/errorcover.png'
		},
        title:{
			type: String,
			default: '无主题视频'
		},
        time:{
            type: null,
			default: '00:00:00'
        },
        watch:{
            type: null,
			      default: '0'
        },
        cid:{
            type: null,
            default: '默认'
        },
        master:{
            type: null,
			      default: 'admin'
        },
    },
    methods:{
    showModal(action) {
        if (action === 'edit') {
            this.modalMessage = '光速开发中，敬请期待';
            this.currentAction = 'edit';  // 标记为编辑操作
        } else if (action === 'delete') {
            this.modalMessage = '确定要删除该视频吗？';
            this.currentAction = 'delete';  // 标记为删除操作
        } else if (action === 'warn'){
            this.modalMessage = '当前权限不足，联系站主获取完整权限~';
            this.currentAction = 'warn';
        }

        this.$refs.confirmModal.show();  // 显示模态框
    },

    confirmAction() {
        if (this.currentAction === 'delete') {
            // 执行删除操作
            this.deleteFunc();
        } else if (this.currentAction === 'edit') {
            // 处理编辑操作，或者显示信息
            console.log('执行编辑操作');
        }else if (this.currentAction === 'warn') {
            // 处理编辑操作，或者显示信息
            console.log('用户权限不足');
        }

    },

    async deleteFunc(){
        let json = {'vid': this.vid}
        let alert = {}
        await postRequest('/admin/delete/video/information',json).then(res => {
            switch (res.code) {
                case 201:
                    alert = {'type':1,'data':'删除途中出现异常'}
                    break
                case 200:
                    alert = {'type':2,'data': '删除该视频成功'}
                    break
                default:
                    alert = {'type':3,'data':res.msg}
            }
        })
        if (alert["type"] === 3){  // 用户权限不足
            // console.log(alert)
            this.showModal('warn')
            return
        }
        this.$refs.AlertMsg.addMsg(
            alert.type,alert.data
        )
        this.$emit('reset');
    },

		onClickVideo(){
        let json ={'vid':this.vid,'uid':this.uid,'titile':this.title,'master':this.master}
        this.$emit('go',json);
    },

    goMaster(){
        let json ={'uid':this.uid,'master':this.master}
        this.$emit('gomaster',json);
    }
	},
}
</script>

<style scoped>

@import "//at.alicdn.com/t/c/font_4674494_dpyktpnc8w.css";
@import "//at.alicdn.com/t/c/font_4674494_c0auq1t19vt.css";

*{
    margin: 0;
    padding: 0;
}


.video-cover-box {
    width: 370px;
    height: 280px;
    background-color: #2D3047;
    box-shadow: 4px 3px 12px #2D3047;
    user-select: none;
    border-radius: 8px;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-bottom: 20px;
    position: relative;
    //border: 3px solid #1C1C1E;
}

.video-cover-box:hover {
    box-shadow: 4px 3px 12px #F3CA40;
    /* background-color: #6A93C3; */
}

.video-cover-box:hover .cover-info {
    width: 100%;
    display: flex;
    animation: slideup 0.3s ease-out;
}

.video-cover-box:hover .video-cover img {
    transform: scale(1.1);
}

.video-cover-box .video-cover {
    background-color: #1C1C1E;
    color: #D3D3D3;
    position: relative;
    overflow: hidden;
    width: 350px;
    height: 190px;
    border-radius: 4px;
    cursor: pointer;
    font-size: 40px;
    font-family: 'SimHei', 'Helvetica', 'Arial', sans-serif; /* 设置为黑体 */
    //display: flex; /* 使用 Flexbox 布局 */
    //justify-content: center; /* 水平方向居中对齐 */
    //align-items: center; /* 垂直方向居中对齐 */
    text-align: center; /* 确保文本在块级元素中居中对齐 */

    border: 3px solid #1C1C1E;
}

.video-cover-box .video-cover img {
    width: 100%; /* 确保图片宽度填满容器 */
    height: 100%; /* 确保图片高度填满容器 */
    object-fit: cover; /* 保持图片比例同时填满容器 */
    transition: 0.5s all ease-in-out;
}

.cover-text {
    display: flex;
    align-items: center; /* 垂直居中 */
    justify-content: center; /* 水平居中 */
    width: 100%;
    height: 100%;
}

.cover-text .iconfont {
    font-size: 50px; /* 图标大小 */
    color: #D3D3D3; /* 图标颜色 */
}

.video-cover img:hover {
    transform: scale(1.1);
}

.cover-info{
    position: absolute;
    //background-color: #FFFFFF;
    //color: #F3CA40;
    width: auto;
    height: 25px;
    right: 0px;
    bottom: 0px;
    display: flex;
    flex-direction: row-reverse;
    align-items: center;
}
.cover-info p{
    padding-right: 5px;
    display: flex;
    flex-direction: row-reverse;
    align-items: center;
    cursor: pointer;
}

.video-cover:hover .cover-info{
    animation: slideup 0.3s ease-out;
    color: #FFFFFF;
}

@keyframes slideup {
    0% {
        height: 0px;
    }
    100% {
        height: 25px;
    }
}
.video-cover-box .video-detail-box{
    margin: 5px 0;
    width: 350px;
    height: 40px;
}

.video-detail-box p{
    //text-overflow: -o-ellipsis-lastline;
	overflow: hidden;
	text-overflow: ellipsis;
	display: -webkit-box;
	-webkit-line-clamp: 2;
	line-clamp: 2;					
	-webkit-box-orient: vertical;
  cursor: pointer;
  width: fit-content;
    font-size: 20px;
    color: #FFFFFF;
}

.video-detail-box p:hover{
    //color: #66a6ff;
}
.video-cover-box .video-info{
    width: 350px;
    height: 20px;
}
.iconColor{
    color: #F3CA40;
    padding-right: 15px;
    font-size: 18px;
}

.iconColor2{
    color: #D3D3D3;
    padding-right: 15px;
    font-size: 24px;
}

.iconColor2:hover{
    color: #F3CA40;
}

.tag {
    opacity: 0.7;
    display: inline-block; /* 确保它像块状一样表现，但保持内联性质 */
    color: #333; /* 标签文字颜色 */
    background-color: #D3D3D3; /* 标签背景颜色 */
    padding: 0px 5px; /* 标签的内边距，调节尺寸 */
    font-size: 18px; /* 文字大小 */
    border-radius: 7px; /* 圆角，使标签看起来更平滑 */
    border: 3px solid #333; /* 边框颜色，略浅，与背景形成对比 */
    margin-right: 15px; /* 为标签之间增加一点间距 */
    //cursor: pointer; /* 鼠标悬停时显示为手型 */
    //box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1); /* 轻微阴影效果，使标签更立体 */
    transition: background-color 0.3s ease; /* 添加交互时的过渡效果 */
}


.iconColor::before{
    padding-right: 5px;
}
.curPointer{
    cursor: pointer;
}
.curPointer:hover{
    color: #66a6ff;
}
</style>
