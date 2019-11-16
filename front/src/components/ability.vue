<template>
	<div class="w-100">
		<div v-if="!update" class="row px-3">
			<span class="col-3 text-left">
				{{selectedAbility.name}} <br>
				<span class="vo">{{selectedAbility.title}}</span>
			</span>
			<span class="col-8 text-left">
				{{selectedAbility.description}}<br>
				<span class="vo">{{vo.description}}</span>
			</span>
			<!-- <span class="col-8 text-left">{{selectedAbility.description || vo.description}}</span> -->
			<span class="col-1 form-inline text-right">
				<button type="submit" class="btn-sm btn-success" @click="update = true">U</button>
				<button type="submit" class="btn-sm btn-danger" @click="$emit('remove')">X</button>
			</span>
		</div>

		<div v-if="update" class="row px-3">
			<v-autocomplete
				v-if="!ability.id"
				:items="items"
				:get-label="getLabel"
				@update-items="updateItems"
				:component-item="template"
				:auto-select-one-item="false"
				@item-selected="selectedItem"
				@input="inputItem"
				placeholder="English Name"
				class="col-2 mt-1"
			></v-autocomplete>
			<input v-model="selectedAbility.name" type="text" class="form-control col-3" placeholder="Translated Name">
			<div class="form-check form-check-inline ml-2 col-5">
				<label class="form-check-label">Type</label>
				<select v-model="type" class="form-control col-4 mx-1">
					<option value="0"></option>
					<option value="1">Magic Ability</option>
					<option value="2">Battle Plan</option>
					<option value="3">Attack Type</option>
				</select>
			</div>
			<textarea v-model="selectedAbility.description" type="text" class="form-control col-10" rows="3" placeholder/>
			<div class="col-2 px-0">
				<button v-if="ability.id || selectedAbility.id" type="submit" class="form-control btn btn-success" @click="save(selectedAbility)">Update</button>
				<button v-if="ability.id" type="submit" class="form-control btn btn-danger" @click="remove(selectedAbility)">Delete</button>
				<button v-if="!ability.id && selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="add(selectedAbility)">Add</button>
				<button v-if="!ability.id && !selectedAbility.id" type="submit" class="form-control btn btn-primary" @click="save(selectedAbility)">Add</button>
			</div>
			<div v-if="selectedAbility.id" class="col-12 text-left vo px-0">{{vo.name}}: {{vo.description}}</div>
		</div>
		<hr>
	</div>
</template>

<script>
import ItemTemplate from "./ItemTemplate.vue";
import Tooltip from "./tooltip.vue";
export default {
	name: "Ability",
	props: ["abilitiesList", "ability"],
	components: { Tooltip },
	watch: {
		ability: function(newVal) {
			this.selectedAbility = newVal;
			if (!this.ability.id){
				this.update=true;
			}
			this.get(this.selectedAbility.id);
		}
	},
	created: function() {
		this.selectedAbility = this.ability;
		if (!this.ability.id){
			this.update=true;
		}
		this.get(this.selectedAbility.id);
	},
	data() {
		return {
			vo: {},
			selectedAbility: {},
			type: 0,
			template: ItemTemplate,
			items: [],
			update: false
		};
	},
	methods: {
		get: function(id) {
			if (id == null) {
				return;
			}
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/abilities/" + id + "/vo")
				.then(function(res) {
					console.log(res);
					this.vo = res.data;
				});
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/abilities/" + id + "?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.selectedAbility = res.data;
				});
		},
		save: function(ability) {
			if (ability.id == null) {
				ability.id = 0;
			}
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/abilities/" + ability.id + "?lang=" + this.$language, ability)
				.then(function(res) {
					console.log(res);
					if (this.ability.id > 0 && this.selectedAbility.id > 0) {
						this.update=false;
					}
					if (res.status === 201) {
						ability.id = res.data;
						this.add(ability);
						this.new(ability);
					} else if (res.status === 200) {
						this.updateAbility();
					}	
				});
		},
		add: function(ability) {
			ability.type = this.type
			this.$emit("add", ability);
		},
		new: function(ability) {
			this.$emit("new", ability);
		},
		updateAbility: function() {
			this.$emit("update");
		},
		remove: function() {
			this.$emit("remove");
		},
		getLabel: function(item) {
			if (!item) {
				return;
			}
			return item.title;
		},
		updateItems(text) {
			this.selectedAbility.title = text
			this.selectedAbility.id = null
			this.items = this.abilitiesList.filter(item =>
				item.title != null
			).filter(item =>
				item.title.toLowerCase().startsWith(text.toLowerCase())
			);
		},
		selectedItem(item) {
			// this.selectedAbility = item;
			this.get(item.id)
		},
		inputItem(item) {
			if (item === null){
				this.type = 0
			}
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
