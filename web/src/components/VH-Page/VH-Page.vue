<!--页面导航-->
<template>
    <div class="page-bar">
        <ul>
            <li v-if="cur>1"><a @click="cur--,pageClick()">上一页</a></li>
    
            <li v-if="cur==1"><a class="banclick">上一页</a></li>
    
            <li v-for="(index,key) in indexs" :key="key" v-bind:class="{ 'active': cur == index}">
                <a @click="btnClick(index)">{{ index }}</a>
            </li>
    
            <li v-if="cur!=all"><a @click="cur++;pageClick()">下一页</a></li>
    
            <li v-if="cur == all"><a class="banclick">下一页</a></li>
    
            <li style="display: flex; align-items: center;margin-left: 30px;float: right;
            padding: 0 2px;font-size: 16px;color: #FFFFFF;line-height: 40px;height: 40px;">
            共<i>{{all}}</i>页，跳转至<input @keyup.enter="goWhere" type="text" onkeyup = "value=value.replace(/[^\d]/g,'')"  v-model="page" />页</li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'VH-Page',
    data () {
        return {
            cur: this.curpage,
            page:null,
        }
    },
    props:{
        all:{
			type: Number,
			default: 1
		},
        curpage:{
			type: Number,
			default: 1
		},
    },
	watch:{
		curpage(newp,oldpa){
			if(newp==1){
				this.cur = newp
			}		
		}	
	},
    methods:{
        goWhere(){
			if(this.page <= this.all ){
				if(this.cur != parseInt(this.page) & parseInt(this.page) > 0 & parseInt(this.page) != null){
                    this.cur = parseInt(this.page)
                    let json={'where':parseInt(this.page)}
                    this.$emit('gowhere',json);
				}else if(parseInt(this.page) == 0){
                    let json={'where':'zero'}
                    this.$emit('gowhere',json);
				}else{
                    let json={'where':'error'}
                    this.$emit('gowhere',json);
				}
				this.page = ''
			}else{
				let json={'where':'nosearch'}
                this.$emit('gowhere',json);
			}
		},
		btnClick(data){//页码点击事件
			if(data != this.cur){
				this.cur = data;
                this.$emit('page',data);
			}
		},
		pageClick(){
			//console.log('现在在'+this.cur+'页');
			this.$emit('page',this.cur);
		},
    },
    computed: {
		indexs(){
			// 定义开始的数字
			var left = 1;
			// 定义结束的数字
			var right = this.all;

			// 存储返回值
			var ar = [];

			// 前提条件：每次都显示condition条页码
			// 最好是个单数
			var condition = 10;

			// 向上取整->就能获取到中间的数字
			var middle = Math.ceil(condition/2);

			// 要分清情况
			//1、在最左边或者在middle的前面
			//2、在中间
			//3、最右边的情况

			// 当总页数超过condition时，需要判断分页情况
			if(this.all>= condition){	
				// 中间的时候
				// 左右都加上（middle-1)
				if(this.cur > middle && this.cur < this.all-(middle-1)){
						left = this.cur - (middle-1)
						right = this.cur + (middle-1)
				}else{
					//最左边或者在condition的中间
					if(this.cur<=middle){
						left = 1
						right = condition
					// 最右边的情况
					// 结束是总条数，开始是condition减去1
					}else{
						right = this.all
						left = this.all -(condition-1)
					}
				}
			}
			while (left <= right){
				ar.push(left)
				left ++
			}
		return ar

	   }
	}
}
</script>

<style  scoped>
*{
    margin: 0;
    padding: 0;
    user-select: none;
}

.page-bar{
	width: 100%;
	display: flex;
	justify-content: center;
}
ul,li{
    margin: 0px;
    padding: 0px;
	  display: flex;
}
li{
    list-style: none
}
.page-bar li:first-child>a {
   margin-left: 0px
}
.page-bar a{
	cursor: pointer;
	outline: none;
	text-align: center;
	border-radius: 4px;
  background-color: #9C89B8;
  border: 3px solid #1C1C1E;
  color: #fff;
	background-image: none;
	transition: all .2s;
	font-size: 16px;
	min-width: 15px;
	margin: 0 2px;
	padding: 0 13px;
	float: left;
	display: block;
	height: 38px;
	line-height: 38px;
}
.page-bar a:hover{
    background: #F3CA40;
    border-color: #2D3047;
    color: #000000;
}
.page-bar a.banclick{
    cursor:not-allowed;
}
.page-bar .active a{
    cursor: default;
    //cursor: pointer;

    background: #F3CA40;
    border-color: #2D3047;
    color: #000000;
}

.page-bar i{
    font-style:normal;
    color: #99a2aa;
    margin: 0px 4px;
    font-size: 16px;
}

.page-bar input{
	padding: 0 10px;
	width: 50px;
	height: 24px;
	line-height: 24px;
	font-size: 16px;
	box-shadow: none;
	color: #2D3047;
  background-color: #FFFFFF;
  border: 3px solid #ccd0d7;
	text-align: center;
	margin: 0 5px;
	border-radius: 4px;
    outline-style: none;
}
.page-bar input:hover{
    //border: 1px solid #00a1d6;
    background-color: #D3D3D3;
    border: 3px solid #ccd0d7;
}

.page-bar input:focus{
    border: 1px solid #F3CA40;
}
</style>
