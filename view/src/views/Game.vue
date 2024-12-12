<template>
    <div @click="MainClick">
        <div v-if="isadmin">
            <h1>Game Page</h1>
            <button @click="AddChallengeShow">添加题目</button>
        </div>
        <div>
            <div><!--首部-->
                <h1>{{ this.game.Name }}</h1>
                <div id="teamhead">
                    <h2>
                        <span>
                            队伍 :
                            {{ this.teamname }}
                        </span>
                        <br>
                        <span>
                            总分:
                            {{ this.team.score }}
                        </span>
                        
                        <span>
                            排名:
                            {{ this.team.rank }}
                        </span>
                    </h2>
                </div>

            </div>
            <div class="divider"></div>

            <div>
                <div class="Class_" v-for="type in this.List">
                    <div v-if="type.length">
                        <h2> {{ type[0].Class }} </h2>
                        <br>
                        <div class="Challenge" v-for="c_ in type" @click="ChallengeClick(c_)" v-if="type.length">
                            <p>{{ c_.Name }}</p>
                            <br>
                            <p>{{ c_.Desc }}</p>
                            <br>
                            <p>
                                score: {{ c_.Score }}
                                <br>
                                <span>
                                    <span>解出: {{ c_.DoneNum }}</span>
                                </span>
                            </p>
                            <div v-if="!this.Done.get(c_.ID)">
                                <p>未解出</p>
                            </div>
                            <div v-if="this.Done.get(c_.ID)">
                                <p>已解出</p>
                            </div>
                        </div>
                        <div class="divider"></div>
                    </div>

                </div>
            </div>

        </div>
        <div v-if="showModala" class="modal">
            <div class="modal-content" @click="KeepModal">

                <p>{{ this.s.Name }}</p>
                <br>
                <p>{{ this.s.Desc }}</p>
                <br>
                <p>score: {{ this.s.Score }}</p>
                <br>
                <div v-if="this.s.Active">
                    <p>
                    <div v-if="!this.Live.get(this.s.ID)"><span><button @click="GetCon(this.s)">创建实例</button></span>
                    </div>
                    <div v-if="this.Live.get(this.s.ID)"><span><button @click="DelCon(this.s)">销毁实例</button></span>
                    </div>
                    <div v-if="this.Live.get(this.s.ID)"><span> 连接: {{ this.Live.get(this.s.ID) }}</span></div>
                    </p>
                </div>

                <div v-if="this.s.FileName != ''">
                    <p>
                        <button @click="downloadFile(this.s)">下载文件</button>
                    </p>
                </div>
                <br>
                <form @submit.prevent="FlagSubmit(flag)">
                    <div>
                        <label for="Flag">Flag:</label>
                        <input type="text" v-model="flag" id="flag" required />
                        <button type="submit">提交</button>
                    </div>
                </form>
                <div v-if="this.isadmin">
                    <button @click="EditChallenge(this.s)">修改</button>
                </div>

            </div>
        </div>
        <div v-if="addModala" class="modal">
            <div class="modal-content" @click="KeepModal">
                <span class="close" @click="closeModal">&times;</span>
                <div id="class_">
                    <label for="class_">Class:</label>
                    <select id="Class_" v-model="addChallengeSelectedClass">
                        <option v-for="option in this.ListName" :key="option" :value="option">
                            {{ option }}
                        </option>
                    </select>
                </div>
                <div id="name">
                    <label for="name">Name:</label>
                    <input type="text" id="Name" v-model="addChallengename" name="name">
                </div>
                <div id="desc">
                    <label for="description">Description:</label>
                    <input type="text" id="Desc" v-model="addChallengedesc" name="desc">
                </div>
                <div id="score">
                    <label for="score">Score:</label>
                    <input type="number" id="Score" v-model="addChallengescore" name="score">
                </div>
                <div id="active">
                    <label for="active">Active:</label>
                    <input type="checkbox" id="Active" v-model="addChallengeactive" name="active"
                        @change="ActiveChangeadd">
                </div>
                <div v-if="this.image_show" id="image">
                    <label for="image">Image:</label>
                    <input type="text" id="Image" v-model="addChallengeimage" name="image">
                </div>
                <div v-if="!this.image_show" id="flag">
                    <label for="flag">Flag:</label>
                    <input type="text" id="Flag" v-model="addChallengeflag" name="flag">
                </div>
                <button type="button" @click="addChallenge">提交</button>
                <button type="button" @click="closeModal">取消</button>
            </div>

        </div>
        <div v-if="editModala" class="modal">
            <div class="modal-content" @click="KeepModal">
                <span class="close" @click="closeModal">&times;</span>
                <div id="class_">
                    <label for="class_">Class:</label>
                    {{ this.s.Class }}

                </div>
                <div id="name">
                    <label for="name">Name:</label>
                    {{ this.s.Name }}
                </div>
                <div id="desc">
                    <label for="description">Description:</label>
                    <textarea id="Desc" v-model="editChallengedesc" name="desc" rows="2" cols="50"></textarea>
                </div>
                <div id="score">
                    <label for="score">MaxScore:</label>
                    <input type="number" id="Score" v-model="editChallengeMaxscore" name="score">
                </div>

                <div v-if="this.image_show" id="image">
                    <div id="active">
                        <label for="active">Active:</label>
                        <input type="checkbox" id="Active" v-model="editChallengeactive" name="active">
                    </div>
                    <label for="image">Image:</label>
                    <nav></nav>
                    <textarea id="Image" v-model="editChallengeimage" name="image" rows="2" cols="50"></textarea>
                    <nav></nav>
                    <label for="image">Port:</label>
                    <input type="number" id="Port" v-model="editChallengeport" name="port">
                </div>
                <div v-if="!this.image_show" id="flag">
                    <label for="flag">Flag:</label>
                    <nav></nav>
                    <textarea type="text" id="Flag" v-model="editChallengeflag" name="flag" rows="2"
                        cols="50"></textarea>
                </div>
                <button type="button" @click="editChallengeSubmit(this.s)">提交</button>
                <button type="button" @click="closeModal">取消</button>
                <div>
                    <button @click="downloadFile(this.s)">下载文件</button>
                    <br>
                    <input type="file" id="file" @change="handleFileUpload" name="file">
                    <button type="button" @click="editChallengeFile(this.s)">提交文件</button>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
import { Url } from '@/config/config';
import { checkAdmin } from '@/logic/check_jwt';
export default {
    name: 'Game',
    data() {
        return {
            id: '',
            teamname: '',
            teamid: '',
            team: {},
            game: {},
            challenge: [],
            Pwn: [],
            Reverse: [],
            Crypto: [],
            Web: [],
            Misc: [],
            List: [],
            ListName: ['PWN', 'REVERSE', 'CRYPTO', 'WEB', 'MISC'],
            isadmin: false,
            showModala: false,
            addModala: false,
            editModala: false,
            isfirstclickmain: false,
            image_show: false,
            s: {},
            Live: new Map(),
            Done: new Map(),
            EditFlags: [],
            BackUrl: Url,
        };
    },
    async created() {
        if (checkAdmin()) {
            this.isadmin = true
        }
        this.id = this.$route.params.id
        this.teamid = this.$route.params.teamid
        this.teamname = this.$route.params.teamname

        await this.ChallengeInfo()
        await this.userInfo()
        await this.GetConInfo()
        console.log(this.Done)
        console.log(this.Live)

    },
    methods: {
        async EditChallenge(c) {
            if (!checkAdmin()) {
                return;
            }
            await this.getChallengeInfoAll(c)
            console.log(c)
            this.editModala = true
            this.editChallengename = c.Name
            this.editChallengedesc = c.Desc
            this.editChallengeMaxscore = c.MaxScore
            this.editChallengeactive = c.Active
            this.image_show = c.Active
            this.editChallengeFilename = c.FileName
            this.editChallengeport = c.Port
            this.editChallengeflag = c.Flags
            this.editChallengeimage = c.ImageName
        },
        AddChallengeShow() {
            this.addModala = true
            this.isfirstclickmain = true
            this.image_show = false
        },
        async ChallengeInfo() {
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            formData.append('game_id', this.id)
            try {
                const response = await fetch(`${Url}/game/show`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取题目信息失败');
                }
                const data = await response.json();
                this.game = data.data.game
                this.challenge = Object.values(data.data.challenges);
                for (let i = 0; i < this.challenge.length; i++) {
                    if (this.challenge[i].Class === 'PWN') {
                        this.Pwn.push(this.challenge[i])
                    } else if (this.challenge[i].Class === 'REVERSE') {
                        this.Reverse.push(this.challenge[i])
                    } else if (this.challenge[i].Class === 'CRYPTO') {
                        this.Crypto.push(this.challenge[i])
                    } else if (this.challenge[i].Class === 'WEB') {
                        this.Web.push(this.challenge[i])
                    } else if (this.challenge[i].Class === 'MISC') {
                        this.Misc.push(this.challenge[i])
                    }
                }
                this.List.push(this.Pwn)
                this.List.push(this.Reverse)
                this.List.push(this.Crypto)
                this.List.push(this.Web)
                this.List.push(this.Misc)
            } catch (error) {
                console.error('Fetch error:', error);
            }
        },
        MainClick() {
            if (!this.isfirstclickmain) {
                this.showModala = false
                this.addModala = false
                this.editModala = false
            }
            this.isfirstclickmain = false
        },
        ChallengeClick(ss) {
            this.s = ss
            this.showModala = true
            this.isfirstclickmain = true
            this.flag = ""
        },
        closeModal() {
            this.showModala = false
            this.addModala = false
            this.editModala = false
        },
        KeepModal() {
            this.isfirstclickmain = true
        },
        async GetConInfo() {
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            formData.append('game_id', this.id)
            formData.append('team_id', this.teamid)
            try {
                const response = await fetch(`${Url}/game/getconinfo`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取连接信息失败');
                }
                const data = await response.json();
                for (let i = 0; i < data.data.length; i++) {
                    this.Live.set(data.data[i].ChallengeID, data.data[i].Nc)
                }

            } catch {
                console.error('Fetch error:', error);
            }
        },
        async GetCon(c) {
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            formData.append('challenge_id', c.ID)
            formData.append('team_id', this.teamid)
            formData.append('game_id', this.id)

            try {
                const response = await fetch(`${Url}/challenge/getcon`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取连接失败');
                }
                const data = await response.json();
                let nc = data.host + ':' + data.port
                this.Live.set(c.ID, nc)
            } catch {
                console.error('Fetch error:', error);
            }
        },
        async DelCon(c) {
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            formData.append('challenge_id', c.ID)
            formData.append('team_id', this.teamid)
            formData.append('game_id', this.id)

            try {
                const response = await fetch(`${Url}/challenge/delcon`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取连接失败');
                }
                const data = await response.json();
                this.Live.delete(s.ID)
            } catch {
                console.error('Fetch error:', error);
            }
        },
        FlagSubmit(flag) {
            const fromData = new FormData()
            fromData.append('name', localStorage.getItem('name'))
            fromData.append('token', localStorage.getItem('jwt'))
            fromData.append('flag', flag)
            fromData.append('challenge_id', this.s.ID)
            fromData.append('team_id', this.teamid)
            fromData.append('game_id', this.id)
            const response = fetch(`${Url}/challenge/flagsubmit`, {
                method: 'POST',
                body: fromData
            });
            if (!response.ok) {
                throw new Error('flag错误');
            }
            const data = response.json();
        },
        async userInfo() {
            const formData = new FormData()
            formData.append('name', localStorage.getItem('name'))
            formData.append('token', localStorage.getItem('jwt'))
            try {

                const formData = new FormData()
                formData.append('name', localStorage.getItem('name'))
                formData.append('token', localStorage.getItem('jwt'))
                formData.append('team_id', this.teamid)
                const response = await fetch(`${Url}/team/info`, {
                    method: 'POST',
                    body: formData
                });
                if (!response.ok) {
                    throw new Error('获取队伍信息失败');
                }
                const data = await response.json();
                this.team = data;
                console.log(this.team)
                for (let i = 0; i < data.challenge.length; i++) {
                    let id = parseInt(data.challenge[i], 10)
                    this.Done.set(id, 1)
                }



            } catch (error) {
                console.error('Fetch error:', error);
            }

        },
        async addChallenge() {
            const formData = new FormData();
            formData.append('name', localStorage.getItem('name'));
            formData.append('token', localStorage.getItem('jwt'));

            formData.append('class', this.addChallengeSelectedClass);
            formData.append('challenge_name', this.addChallengename);
            formData.append('active', this.addChallengeactive);
            formData.append('max_score', this.addChallengescore);
            formData.append('image_name', this.addChallengeimage);
            if (!this.addChallengeflag) {
                this.addChallengeflag = 'flag{[[FLAG]]}';
            }
            formData.append('flags', this.addChallengeflag);
            formData.append('desc', this.addChallengedesc);
            formData.append('game_id', this.id);
            const res = await fetch(Url + '/challenge/addchallenge', {
                method: 'POST',
                body: formData
            });
            const data = await res.json();
            if (data.status != 200) {
                alert(data.message);
                return;
            }
            alert('添加成功');

        },
        ActiveChangeadd() {
            this.image_show = this.addChallengeactive;
            if (this.active) {
                this.flag = 'flag{[[FLAG]]}';
            } else {
                if (this.flag == 'flag{[[FLAG]]}')
                    this.flag = '';

            }
        },
        IsFileChange() {
            this.file_show = this.is_file;
        },
        async downloadFile(c) {
            const fromData = new FormData();
            fromData.append('name', localStorage.getItem('name'));
            fromData.append('token', localStorage.getItem('jwt'));
            fromData.append('challenge_id', c.ID);

            try {
                const response = await fetch(Url + '/challenge/getfile', {
                    method: 'POST',
                    body: fromData,
                });

                if (!response.ok) {
                    throw new Error('网络请求失败，状态码: ' + response.statusText);
                }

                const blob = await response.blob();

                // 使用 this.s.FileName 作为文件名
                const fileName = this.s?.FileName || "默认文件名";

                // 创建一个隐藏的 a 标签用于触发下载
                const downloadElement = document.createElement("a");
                const href = window.URL.createObjectURL(blob);
                downloadElement.href = href;
                downloadElement.download = c.FileName
                downloadElement.style.display = "none";

                // 将元素添加到页面并触发点击事件
                document.body.appendChild(downloadElement);
                downloadElement.click();
                document.body.removeChild(downloadElement);

                // 释放 URL
                window.URL.revokeObjectURL(href);
            } catch (error) {
                console.error('下载出错:', error);
            }
        }


        ,
        handleFileUpload(event) {
            this.file = event.target.files[0];
        },
        async editChallengeFile(c) {
            const formData = new FormData();
            formData.append('name', localStorage.getItem('name'));
            formData.append('token', localStorage.getItem('jwt'));
            formData.append('challenge_id', c.ID);
            formData.append('file', this.file);
            const res = await fetch(Url + '/challenge/updatefile', {
                method: 'POST',
                body: formData
            });
            const data = await res.json();
            if (res.status != 200) {
                alert(data.message);
                return;
            }
            alert('修改成功');
        },
        async editChallengeSubmit(c) {
            const formData = new FormData();
            formData.append('name', localStorage.getItem('name'));
            formData.append('token', localStorage.getItem('jwt'));
            formData.append('challenge_id', c.ID);
            formData.append('desc', this.editChallengedesc);
            formData.append('max_score', this.editChallengeMaxscore);
            formData.append('image_name', this.editChallengeimage);
            formData.append('flags', this.editChallengeflag);
            formData.append('port', this.editChallengeport);
            const res = await fetch(Url + '/challenge/updateinfo', {
                method: 'POST',
                body: formData
            });
            const data = await res.json();
            if (res.status != 200) {
                alert(data.message);
                return;
            }
            alert('修改成功');
        },
        async getChallengeInfoAll(c) {
            const formData = new FormData();
            formData.append('name', localStorage.getItem('name'));
            formData.append('token', localStorage.getItem('jwt'));
            formData.append('challenge_id', c.ID);
            const res = await fetch(Url + '/challenge/getallchallengeinfo', {
                method: 'POST',
                body: formData
            });
            const data = await res.json();
            if (res.status != 200) {
                alert(data.message);
                return
            }
            c.Flags = data.data.flags;
            c.ImageName = data.data.image;
        }
    },
}

</script>

<style>
.Class_ {
    margin-bottom: 20px;
}

.Challenge {
    padding: 10px;
    margin: 10px;
    width: 200px;
    height: 200px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
}

.divider {
    height: 2px;
    background-color: #eee;
    margin: 20px 0;
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
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.5);
}

.modal-content {
    background-color: #fefefe;
    margin: auto;
    padding: 20px;
    border: 1px solid #888;
    width: 80%;
    max-width: 500px;
    border-radius: 10px;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
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