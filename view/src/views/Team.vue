<template>
    <div class="section" id="section1">
        <h1>Team Page</h1>
        <button @click="addTeam" type="button">创建队伍</button>
        <button @click="joinTeam" type="button">加入队伍</button>
    </div>
    <div class="divider"></div>
    <div class="section" id="section2" >
        <h2>队伍信息</h2>
        <br>
      <ul>
        <div class="team-list">
        <div v-for="team_ in team" :key="team.id" id="team_info">
          <p>队伍名称 {{ team_.name }}</p>
          <p>队伍描述 {{ team_.desc }}</p>
            <p>队伍密钥 {{ team_.key }}</p>
            <p>队长 {{ team_.leader }}</p> 
            <div v-if="team_.members">
                <div v-for="member_ in team_.members">
                    <p>队员 {{ member_ }}</p>
                </div>
            </div>
            <button type="button">修改信息</button>
            <button type="button">退出队伍</button>
            <button type="button">踢出成员</button>
        </div>
    </div>
      </ul>
    </div>
    <div v-if="showModaladd" class="modal">
            <div class="modal-content">
                <span class="close" @click="closeModal">&times;</span>
                <h2>添加队伍</h2>
                <form @submit.prevent="addTeamSubmit">
                <div>
                    <label for="teamName">队伍名称:</label>
                    <input type="text" v-model="newTeamName" id="teamName" required />
                </div>
                <div>
                    <label for="teamDesc">队伍描述:</label>
                    <input type="text" v-model="newTeamDesc" id="teamDesc" required />
                </div>
                <button  type="submit">提交</button>
                <button type="button" @click="closeModal">取消</button>
                </form>
            </div>
        </div>
        <div v-if="showModaljoin" class="modal">
            <div class="modal-content">
                <span class="close" @click="closeModal">&times;</span>
                <h2>加入队伍</h2>
                <form @submit.prevent="joinTeamSubmit">
                <div>
                    <label for="teamName">队伍名称:</label>
                    <input type="text" v-model="JoinTeamName" id="teamName" required />
                </div>
                <div>
                    <label for="teamDesc">队伍key:</label>
                    <input type="text" v-model="JoinTeamKey" id="teamDesc" required />
                </div>
                <button  type="submit">提交</button>
                <button type="button" @click="closeModal">取消</button>
                </form>
            </div>
        </div>
</template>

<script>
import { Url } from '@/config/config';
export default {
    name : 'Team',
    data() {
        return{
            team: [],
            newTeam: {},
            showModaladd: false, 
            showModaljoin: false,
        };
    },
    async created(){
        this.userInfo()
    },
    methods: {
        async userInfo() {
            const formData = new FormData()
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            try {
                const response = await fetch(`${Url}/user/info`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取信息失败');
                }
                const data = await response.json();
                const teamId = data.Team.split(',').map(id => parseInt(id.trim(), 10)).filter(Number.isInteger);
                for(let i=0;i<teamId.length;i++){
                    const formData = new FormData()
                    formData.append('name',localStorage.getItem('name'))
                    formData.append('token',localStorage.getItem('jwt'))
                    formData.append('team_id',teamId[i])
                    try {
                        const response = await fetch(`${Url}/team/info`, {
                            method: 'POST',
                            body: formData
                        });
                        if (!response.ok) {
                            throw new Error('获取队伍信息失败');
                        }
                        const data = await response.json();
                        this.team.push(data)
                        console.log(this.team[i])
                    }catch(error){
                        console.error('Fetch error:', error);
                    }
                }
            }catch(error){
                console.error('Fetch error:', error);
            }

        },
        addTeam(){
            this.showModaladd=true
        },
        async addTeamSubmit() {
            const formData = new FormData()
            formData.append('team',this.newTeamName)
            formData.append('desc',this.newTeamDesc)
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            try {
                const response = await fetch(`${Url}/team/create`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('添加队伍失败');
                }
            }catch(error){
                console.error('Fetch error:', error);
            }
            this.showModaladd=false
        },
        joinTeam(){
            this.showModaljoin=true
        },
        async joinTeamSubmit(){
            const formData = new FormData()
            formData.append('team',this.JoinTeamName)
            formData.append('key',this.JoinTeamKey)
            formData.append('name',localStorage.getItem('name'))
            formData.append('token',localStorage.getItem('jwt'))
            try {
                const response = await fetch(`${Url}/team/join`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('添加队伍失败');
                }
            }catch(error){
                console.error('Fetch error:', error);
            }
            this.showModaljoin=false
        },
        closeModal(){
            this.showModaladd=false
            this.showModaljoin=false
        },
    }
}

</script>

<style>
.team-list {
  display: flex;
  flex-wrap: wrap;
  gap: 20px; /* 控制队伍信息之间的间距 */
}
#team_info {
  padding: 20px;
  margin: 10px 0;
  width: 300px;
  border: 1px solid #ccc; /* 添加边框 */
  border-radius: 5px; /* 添加圆角 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}
.section {
  padding: 20px;
  margin: 10px 0;
}

.divider {
  height: 2px; /* 分隔线的高度 */
  background-color: #ccc; /* 分隔线的颜色 */
  margin: 20px 0; /* 分隔线的上下间距 */
}

.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 50%;
  height: 50%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.5);
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
</style>