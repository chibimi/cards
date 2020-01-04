<template>
	<div class="spell">
		<div v-if="!update" class="row">
			<span class="col-2">
				{{ spell.name }} <br />
				<span class="vo">{{ spell.title }}</span>
			</span>
			<span class="col-2">
				Cost: {{ spell.cost }}, Range: {{ spell.rng }}<br />
				AoE: {{ spell.aoe }}, Pow: {{ spell.pow }}<br />
				Dur: {{ spell.dur }}, Off: {{ spell.off }}
			</span>
			<span class="col-7">
				{{ spell.description }}<br />
				<span class="vo">{{ vo.description }}</span>
			</span>
			<span class="col-1">
				<button class="btn-success mb-1" @click="startUpdate()">Update</button>
				<button class="btn-danger" @click="$emit('remove')">Delete</button>
			</span>
		</div>

		<div v-if="update" class="row">
			<span v-if="newSpell" class="col-2 ">English Name</span>
			<span class="col-2 ">Name</span>
			<span class="col-1 ">Cost</span>
			<span class="col-1 ">Range</span>
			<span class="col-1 ">Aoe</span>
			<span class="col-1 ">Pow</span>
			<span class="col-1 ">Dur</span>
			<span class="col-1 ">Off</span>
			<span class="col-2"></span>
			<span v-if="!newSpell" class="col-2"></span>

			<v-autocomplete
				v-if="newSpell"
				:items="items"
				:get-label="getLabel"
				@update-items="updateItems"
				:component-item="template"
				:auto-select-one-item="false"
				@item-selected="selectedItem"
				@input="inputItem"
				placeholder="English Name"
				class="col-2 pr-0"
			></v-autocomplete>

			<input v-model="spell.name" class="col-2" :class="{ 'ml-3': !newSpell }" placeholder="Translated Name" />
			<input v-model="spell.cost" class="col-1" placeholder="cost" />
			<input v-model="spell.rng" class="col-1" placeholder="rng" />
			<input v-model="spell.aoe" class="col-1" placeholder="aoe" />
			<input v-model="spell.pow" class="col-1" placeholder="pow" />
			<input v-model="spell.dur" class="col-1" placeholder="dur" />
			<input v-model="spell.off" class="col-1" placeholder="off" />
			<span class="col-2 py-1 text-danger">USE ENGLIGH HERE</span>
			<span v-if="!newSpell" class="col-2"></span>
			<div class="col-11">
				<TextArea
					v-model="spell.description"
					:ref_id="spell_id"
					:abilities="abilities"
					class="w-100"
					:rows="3"
					placeholder="Translated spell description"
				/>
			</div>

			<div class="col-1 pl-0">
				<button v-if="spell.id && !newSpell" class="btn-success" @click="save(spell)">Update</button>
				<button v-if="spell.id && !newSpell" class="my-1" @click="cancelUpdate()">Cancel</button>
				<button v-if="newSpell" @click="save(spell)">Add</button>
			</div>

			<div v-if="spell.id" class="col-12 vo">{{ vo.name }}: {{ vo.description }}</div>
		</div>
		<hr />
	</div>
</template>

<script>
import ItemTemplate from './ItemTemplate.vue'
import TextArea from './textarea.vue'

export default {
	name: 'Spell',
	props: ['spellsList', 'spell_id', 'abilities'],
	components: { TextArea },
	watch: {},
	created: function() {
		if (!this.spell_id) {
			this.update = true
			this.newSpell = true
		}
		this.get(this.spell_id)
	},
	data() {
		return {
			spell: {},
			vo: {},
			template: ItemTemplate,
			items: [],
			update: false,
			beforeEdit: {},
			newSpell: false,
		}
	},
	methods: {
		startUpdate:function(){
			this.update=true
			this.beforeEdit = JSON.parse(JSON.stringify(this.spell))
		},
		cancelUpdate:function(){
			this.update=false
			this.spell = JSON.parse(JSON.stringify(this.beforeEdit))
		},
		get: function(id) {
			if (id == null) {
				return
			}
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/spells/${id}?lang=US`)
				.then(function(res) {
					console.debug(res)
					this.vo = res.data
				})
				.catch(function(err) {
					console.error(err)
				})
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/spells/${id}?lang=${this.$language}`)
				.then(function(res) {
					console.log(res)
					this.spell = res.data
				})
				.catch(function(err) {
					console.error(err)
				})
		},
		save: function(spell) {
			if (spell.id == null) {
				spell.id = 0
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + `/spells/${spell.id}?lang=${this.$language}`, spell)
				.then(function(res) {
					console.debug(res)
					if (res.status === 201) {
						spell.id = res.data
					}
					if (!this.newSpell) {
						this.update = false
					} else {
						this.$emit('add', spell)
						this.spell = {}
					}
					this.$emit('update')
				})
				.catch(function(err) {
					console.error(err)
				})
		},

		// Handle Autocomplete
		getLabel: function(item) {
			if (!item) {
				return
			}
			return item.title
		},
		updateItems(text) {
			this.spell.title = text
			this.spell.id = null
			this.items = this.spellsList
				.filter(item => item.title != null)
				.filter(item => item.title.toLowerCase().startsWith(text.toLowerCase()))
		},
		selectedItem(item) {
			this.get(item.id)
		},
		inputItem(item) {
			if (item === null) {
				this.spell = {}
			}
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
.spell {
	button {
		@extend .btn-sm;
	}
}
</style>
