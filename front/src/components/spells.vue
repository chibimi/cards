<template>
	<div class="w-100  mt-4">
		<h4 class="text-left">Spells</h4>
		<Spell
			v-for="(value,index) in spells"
			v-bind:spell="value"
			:spellsList="spellsList"
			:key="value.id"
			v-on:remove="removeSpell(value,index)"
			v-on:update="updateSpell"
		></Spell>
		<div class="card border-secondary">
			<h5
				class="card-header bg-secondary text-light card-icon py-1"
				data-toggle="collapse"
				data-target="#new_card_spell"
				aria-expanded="false"
				aria-controls="new_card_spell"
			>New Spell</h5>
			<div class="collapse card-body p-1" id="new_card_spell">
				<Spell :spell="spell" :spellsList="spellsList" v-on:add="addSpell" v-on:new="newSpell" v-on:update="updateSpell"></Spell>
			</div>
		</div>
		<hr>
	</div>
</template>

<script>
import Spell from "./spell.vue";
export default {
	name: "Spells",
	props: ["id"],
	components: {
		Spell
	},
	watch: {
		id: function(newVal) {
			this.get(newVal);
		}
	},
	created: function() {
		this.get(this.id);
		this.getSpells();
	},
	data() {
		return {
			spellsList: [],
			spells: [],
			spell: {}
		};
	},
	methods: {
		get: function(cardID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + cardID + "/spell?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.spells = res.data;
				});
		},
		getSpells: function() {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/spells?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.spellsList = res.data;
				});
		},
		removeSpell: function(spell, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + this.id + "/spell/" + spell.id)
				.then(function(res) {
					console.log(res);
					if (res.status === 204) {
						this.spells.splice(index, 1);
					}
				});
		},
		addSpell: function(spell) {
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT+ "/ref/" + this.id + "/spell/" + spell.id + "?lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					if (res.status === 201) {
						this.spells.push(spell);
						this.spell = {};
					}
				});
		},
		updateSpell: function() {
			this.getSpells();

		},
		newSpell: function(spell) {
			this.spellsList.push(spell);
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
