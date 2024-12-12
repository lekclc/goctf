<template>
    <div @click="MainClick">
        <div>
            <h1>GameList Page</h1>
            <div v-if="isadmin">
                <button type="button" @click="addgame">添加比赛</button>
            </div>
            <div class="divider"></div>
            <div v-for="game in gameslist" :key="game.id" class="game">
                <p>比赛名称 {{ game.Name }}</p>
                <p>比赛描述 {{ game.Desc }}</p>
                <p>比赛开始时间 {{ game.Start }}</p>
                <p>比赛结束时间 {{ game.End }}</p>
                <button type="button" @click="JmpGame(game)">进入比赛</button>

                <button type="button" @click="JmpRank(game)">查看榜单</button>
                <button type="button" @click="SignUpFirst(game)" v-if="!game.team[0]">报名比赛</button>
                <span v-if="game.team[0]"> 已经报名 {{this.team_id_name[game.team[0]] }}</span>
                <div v-if="isadmin">
                    <button type="button">修改信息</button>
                </div>

            </div>

        </div>
        <div v-if="addgamemodel" class="modal">
            <div class="modal-content" @click="ModalClick">
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
                    <button type="submit">提交</button>
                    <button type="button" @click="closeModal">取消</button>
                </form>
            </div>
        </div>
        <div v-if="signupgamemodel" class="modal">
            <div class="modal-content" @click="ModalClick">
                <span class="close" @click="closeModal">&times;</span>
                <h2>报名比赛</h2>
                <form @submit.prevent="signupsubmit">
                    <div>
                        <label for="teamSelect">选择队伍:</label>
                        <select id="teamSelect" v-model="selectedTeam" required>
                            <option v-for="team_ in activeItemList" :key="team_.id" :value="team_.id">
                                {{ team_.name }}
                            </option>
                        </select>
                    </div>
                    <button type="submit">提交</button>
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
            isadmin: false,
            signupgamemodel: false,
            isfirst: true,
            game_team: [],
            team: [],
            joingameid: '',
            name: '',
            team_id_name: {},
        };
    },
    computed: {
            activeItemList: function () {
                return this.team.filter((item) => {
                    return item.gameID === 0 && item.leader === this.name;
                })
            },
        },
    async created() {

        if (checkAdmin()) {
            this.isadmin = true;
        }
        this.name = localStorage.getItem('name');
        await this.GetGameList();
        await this.userInfo();

    },
    methods: {

        SignUpFirst(game) {
            this.signupgamemodel = true;
            this.isfirst = true
            this.joingameid = game.ID;
        },
        async signupsubmit() {
            console.log(this.selectedTeam, this.joingameid)
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            formData.append('team_id', this.selectedTeam)
            formData.append('game_id', this.joingameid)
            try {
                const response = await fetch(`${Url}/game/join`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('报名比赛失败');
                }
                const data = await response.json();
            } catch (error) {
                console.error(error);
            }
            this.signupgamemodel = false;
        },
        async GetGameList() {
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
                for (let i = 0; i < this.gameslist.length; i++) {
                    this.gameslist[i].Start = new Date(this.gameslist[i].Start).toLocaleString();
                    this.gameslist[i].End = new Date(this.gameslist[i].End).toLocaleString();
                    this.game_team.push({
                        game: this.gameslist[i].ID,
                        team: [],
                    });
                }
            } catch (error) {
                console.error(error);
            }
        },
        async userInfo() {
            const formData = new FormData();
            formData.append('name', localStorage.getItem('name'));
            formData.append('token', localStorage.getItem('jwt'));
            try {
                const response = await fetch(`${Url}/user/info`, {
                    method: 'POST',
                    body: formData,
                });
                if (!response.ok) {
                    throw new Error('获取信息失败');
                }
                const data = await response.json();
                const teamId = data.Team.split(',').map(id => parseInt(id.trim(), 10)).filter(Number.isInteger);
                for (let i = 0; i < teamId.length; i++) {
                    const formData = new FormData();
                    formData.append('name', localStorage.getItem('name'));
                    formData.append('token', localStorage.getItem('jwt'));
                    formData.append('team_id', teamId[i]);

                    try {
                        const response = await fetch(`${Url}/team/info`, {
                            method: 'POST',
                            body: formData,
                        });
                        if (!response.ok) {
                            throw new Error('获取队伍信息失败');
                        }
                        const teamData = await response.json();
                        this.team.push(teamData);
                        this.team_id_name[teamData.id] = teamData.name;
                        for (let j = 0; j < this.game_team.length; j++) {
                            if (this.game_team[j].game === this.team[i].gameID) {
                                this.game_team[j].team.push(teamId[i]);
                            }
                            this.gameslist[j].team = this.game_team[j].team;
                        }

                    } catch (error) {
                        console.error('Fetch error:', error);
                    }
                }
            } catch (error) {
                console.error('Fetch error:', error);
            }
        },
        async addgamesubmit() {
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            formData.append('game', this.SetGameName)
            formData.append('desc', this.SetGameDesc)
            formData.append('start', this.SetStartTime)
            formData.append('end', this.SetEndTime)
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
            this.addgamemodel = false;
        },
        MainClick() {
            if (this.isfirst) {
                this.isfirst = false
            } else {
                this.isfirst = true
                this.addgamemodel = false;
                this.signupgamemodel = false;
            }
        },
        ModalClick() {
            this.isfirst = true
        },
        addgame() {
            this.addgamemodel = true;
        },
        closeModal() {
            this.addgamemodel = false;
            this.signupgamemodel = false;
        },
        JmpGame(game) {
            this.$router.push({ name: 'Game', params: { id: game.ID,teamid: game.team[0] ,teamname: this.team_id_name[game.team[0]] } });
        },
        JmpRank(game){
            this.$router.push({ name: 'Rank', params: { id: game.ID,page: 1 } });
        }
    },
}
</script>

<style>
.modal {
    display: block;
    /* 显示模态框 */
    position: fixed;
    /* 固定定位 */
    z-index: 1;
    /* 在最上层 */
    left: 0;
    top: 0;
    width: 100%;
    /* 全屏 */
    height: 100%;
    /* 全屏 */
    background-color: rgb(0, 0, 0);
    /* 背景颜色 */
    background-color: rgba(0, 0, 0, 0.4);
    /* 背景颜色和透明度 */

}

.modal-content {
    background-color: #fefefe;
    margin: 15% auto;
    /* 15% 从顶部和居中 */
    padding: 20px;
    border: 1px solid #888;
    width: 80%;
    /* 宽度 */
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
    height: 2px;
    /* 分隔线的高度 */
    background-color: #ccc;
    /* 分隔线的颜色 */
    margin: 20px 0;
    /* 分隔线的上下间距 */
}

.game {
    padding: 20px;
    margin: 10px 0;
    width: 100%;
    border: 1px solid #ccc;
    /* 添加边框 */
    border-radius: 5px;
    /* 添加圆角 */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    /* 添加阴影 */
}
</style>