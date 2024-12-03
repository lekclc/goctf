<template>
    <div>
        <h1>GameList Page</h1>
        <div v-if="isadmin">
            <button type="button" @click="addgame">添加比赛</button>
        </div>
        <div class="divider"></div>
        <div v-for="game in gameslist" :key="game.id" class="game" >
            <p>比赛名称 {{ game.Name }}</p>
            <p>比赛描述 {{ game.Desc }}</p>
            <p>比赛开始时间 {{ game.Start }}</p>
            <p>比赛结束时间 {{ game.End }}</p>
            <button type="button" @click="JmpGame(game.ID)">进入比赛</button>

            <button type="button">查看榜单</button>
            <div v-if="isadmin">
            <button type="button">修改信息</button>
            </div>

        </div>
        <div v-if="addgamemodel" class="modal">
            <div class="modal-content">
                <span class="close" @click="closeModal">&times;</span>
                <h2>加入队伍</h2>
                <form @submit.prevent="addgamesubmit">
                <div>
                    <label for="gameName">比赛名称:</label>
                    <input type="text" v-model="SetGameName" id="gameName" required />
                </div>
                <div>
                    <label for="teamDesc">比赛描述:</label>
                    <input type="text" v-model="SetGameDesc" id="gameDesc" required />
                </div>
                <div>
                    <label for="gameStartTime">比赛开始时间:</label>
                    <input type="datetime-local" v-model="SetStartTime" id="gameStartTime" required />
                </div>
                <div>
                    <label for="gameEndTime">比赛结束时间:</label>
                    <input type="datetime-local" v-model="SetEndTime" id="gameEndTime" required />
                </div>
                <button  type="submit">提交</button>
                <button type="button" @click="closeModal">取消</button>
                </form>
            </div>
        </div>
    </div>
    
</template>

<script>
import { Url } from '@/config/config';
import { checkAdmin } from '@/logic/check_jwt';
export default {
    name: 'GameList',
    data() {
        return {
            gameslist: [],
            addgamemodel: false,
            isadmin: false
        };
    },
    async created() {
        this.GetGameList();
        if(checkAdmin()){
            this.isadmin = true;
        }
    },
    methods: {
        async GetGameList(){
            try {
                const response = await fetch(`${Url}/game/gamelist`, {
                    method: 'POST',
                });
                if (!response.ok) {
                    throw new Error('获取比赛信息失败');
                }
                const data = await response.json();
                this.gameslist = Object.values(data.data);
                console.log(this.gameslist)
            } catch (error) {
                console.error(error);
            }
        },
        async addgamesubmit() {
            const formData = new FormData()
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            formData.append('game',this.SetGameName)
            formData.append('desc',this.SetGameDesc)
            formData.append('start',this.SetStartTime)
            formData.append('end',this.SetEndTime)
            try {
                const response = await fetch(`${Url}/game/set`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('添加比赛失败');
                }
                const data = await response.json();
            } catch (error) {
                console.error(error);
            }
            this.addgamemodel=false;
        },
        addgame() {
            this.addgamemodel = true;
        },
        closeModal() {
            this.addgamemodel = false;
        },
        JmpGame(id) {
            this.$router.push(`/game/${id}`);
        },
    },
}
</script>

<style>

.modal {
  display: block; /* 显示模态框 */
  position: fixed; /* 固定定位 */
  z-index: 1; /* 在最上层 */
  left: 0;
  top: 0;
  width: 100%; /* 全屏 */
  height: 100%; /* 全屏 */
  background-color: rgb(0,0,0); /* 背景颜色 */
  background-color: rgba(0,0,0,0.4); /* 背景颜色和透明度 */
  
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto; /* 15% 从顶部和居中 */
  padding: 20px;
  border: 1px solid #888;
  width: 80%; /* 宽度 */
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

.divider {
  height: 2px; /* 分隔线的高度 */
  background-color: #ccc; /* 分隔线的颜色 */
  margin: 20px 0; /* 分隔线的上下间距 */
}

.game {
    padding: 20px;
    margin: 10px 0;
    width: 100%;
    border: 1px solid #ccc; /* 添加边框 */
    border-radius: 5px; /* 添加圆角 */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}
</style>